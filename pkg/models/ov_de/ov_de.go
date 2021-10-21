// Package ovde implements Ownershipvoucher deserialization from CBOR
// As for our needs we'll deserialize its header only
package ovde

import (
	"bytes"
	"fmt"
	"github.com/fxamacker/cbor/v2"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

// Extract FDO uuid from the OV's header to a valid uuid string
// Panic if can't
func guidAsString(ovh *OwnershipVoucherHeader) string {
	return fmt.Sprint(uuid.Must(uuid.FromBytes(ovh.GUID)))
}

// Extract device name from the OV's header
func deviceName(ovh *OwnershipVoucherHeader) string {
	return ovh.DeviceInfo
}

// Extract device protocol version from the OV's header
func deviceProtocolVersion(ovh *OwnershipVoucherHeader) uint16 {
	return ovh.ProtocolVersion
}

// CBOR unmarshal of OV header, receives []byte from unmarshalOwnershipVoucher
// returns OV header as pointer to OwnershipVoucherHeader struct & err
func unmarshalOwnershipVoucherHeader(ovhb []byte) (*OwnershipVoucherHeader, error) {
	var ovh OwnershipVoucherHeader
	err := cbor.Unmarshal(ovhb, &ovh)
	return &ovh, err
}

// If CBOR unmarshal fails => panic
// Something might be wrong with OV
func unmarshalCheck(e error, ovORovh string, ovNum int) {
	if e != nil {
		log.WithFields(log.Fields{
			"method": "ovde.unmarshalCheck",
			"what":   ovORovh,
			"ov_num": ovNum,
		}).Error(e)
		panic(e)
	}
}

// CBOR unmarshal of OV, receives []byte from loading the OV file (either reading/receiving)
// do some validation checks and returns OV header as pointer to OwnershipVoucherHeader struct
func parseBytes(ovb []byte) []OwnershipVoucherHeader {
	var (
		ov        OwnershipVoucher
		ovh       []OwnershipVoucherHeader
		counter   int        = 0
		logFields log.Fields = map[string]interface{}{"method": "ovde.parseBytes"}
	)
	defer func() { // in a panic case, stop the parsing but keep alive
		if err := recover(); err != nil {
			logFields["ovs_parsed"] = counter
			log.WithFields(logFields).Error("panic occurred")
		}
	}()
	if err := cbor.Valid(ovb); err == nil { // checking whether the CBOR data is complete and well-formed
		dec := cbor.NewDecoder(bytes.NewReader(ovb))
		for { // stream OVs
			if decErr := dec.Decode(&ov); decErr == io.EOF {
				break
			} else if decErr != nil { // couldn't decode into ownershipvoucher
				unmarshalCheck(decErr, "ownershipvoucher", counter)
			}
			singleOvh, err := unmarshalOwnershipVoucherHeader(ov.Header)
			unmarshalCheck(err, "ownershipvoucher header", counter)
			ovh = append(ovh, *singleOvh)
			counter++
		}
		if counter > 0 && dec.NumBytesRead() == len(ovb) {
			logFields["ovs_parsed"] = counter
			log.WithFields(logFields).Info("All ownershipvoucher bytes parsed successfully")
		}
	} else {
		logFields["ovs_parsed"] = counter
		log.WithFields(logFields).Error("Invalid ownershipvoucher bytes")
	}
	return ovh
}

// Get minimum data required from parseBytes without marshal the whole OV header to JSON (though possible)
func minimumParse(ovb []byte) []map[string]interface{} {
	ovh := parseBytes(ovb)
	var minimumDataReq []map[string]interface{}
	for _, header := range ovh {
		data := map[string]interface{}{
			"device_name":      deviceName(&header),
			"fdo_uuid":         guidAsString(&header),
			"protocol_version": deviceProtocolVersion(&header),
		}
		minimumDataReq = append(minimumDataReq, data)
		data["method"] = "ovde.minimumParse"
		log.WithFields(data).Debug("New device added")
	}
	return minimumDataReq
}
