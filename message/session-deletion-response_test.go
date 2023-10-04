// Copyright 2019-2022 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package message_test

import (
	"testing"
	"time"

	"github.com/wmnsk/go-pfcp/ie"
	"github.com/wmnsk/go-pfcp/message"

	"github.com/wmnsk/go-pfcp/internal/testutil"
)

func TestSessionDeletionResponse(t *testing.T) {
	cases := []testutil.TestCase{
		{
			Description: "Single IE",
			Structured: message.NewSessionDeletionResponse(
				mp, fo, seid, seq, pri,
				ie.NewCause(ie.CauseRequestAccepted),
				ie.NewOffendingIE(ie.Cause),
				ie.NewLoadControlInformation(ie.NewSequenceNumber(0xffffffff), ie.NewMetric(0x01)),
				ie.NewOverloadControlInformation(
					ie.NewSequenceNumber(0xffffffff),
					ie.NewMetric(0x01),
					ie.NewTimer(20*time.Hour),
					ie.NewOCIFlags(0x01),
				),
				ie.NewUsageReportWithinSessionDeletionResponse(
					ie.NewURRID(0xffffffff),
					ie.NewURSEQN(0xffffffff),
					ie.NewUsageReportTrigger(0xff, 0xff, 0xff),
					ie.NewStartTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewEndTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewVolumeMeasurement(0x3f, 0x1111111111111111, 0x2222222222222222, 0x3333333333333333, 0x4444444444444444, 0x5555555555555555, 0x6666666666666666),
					ie.NewDurationMeasurement(10*time.Second),
					ie.NewTimeOfFirstPacket(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewTimeOfLastPacket(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewUsageInformation(1, 1, 1, 1),
					ie.NewEthernetTrafficInformation(
						ie.NewMACAddressesDetected(
							ie.NewCTAG(0x07, 1, 1, 4095),
							ie.NewSTAG(0x07, 1, 1, 4095),
							mac1, mac2, mac3, mac4,
						),
						ie.NewMACAddressesRemoved(
							ie.NewCTAG(0x07, 1, 1, 4095),
							ie.NewSTAG(0x07, 1, 1, 4095),
							mac1, mac2, mac3, mac4,
						),
					),
				),
				ie.NewAdditionalUsageReportsInformation(0x00ff),
			),
			Serialized: []byte{
				0x21, 0x37, 0x01, 0x14, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x11, 0x22, 0x33, 0x00,
				0x00, 0x13, 0x00, 0x01, 0x01,
				0x00, 0x28, 0x00, 0x02, 0x00, 0x13,
				0x00, 0x33, 0x00, 0x0d, 0x00, 0x34, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff, 0x00, 0x35, 0x00, 0x01, 0x01,
				0x00, 0x36, 0x00, 0x17,
				0x00, 0x34, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0x35, 0x00, 0x01, 0x01,
				0x00, 0x37, 0x00, 0x01, 0x82,
				0x00, 0x6e, 0x00, 0x01, 0x01,
				0x00, 0x4f, 0x00, 0xc7,
				0x00, 0x51, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0x68, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0x3f, 0x00, 0x03, 0xff, 0xff, 0xff,
				0x00, 0x4b, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x4c, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x42, 0x00, 0x31,
				0x3f,
				0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11,
				0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22,
				0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33,
				0x44, 0x44, 0x44, 0x44, 0x44, 0x44, 0x44, 0x44,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
				0x00, 0x43, 0x00, 0x04, 0x00, 0x00, 0x00, 0x0a,
				0x00, 0x45, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x46, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x5a, 0x00, 0x01, 0x0f,
				0x00, 0x8f, 0x00, 0x4a,
				0x00, 0x90, 0x00, 0x21,
				0x04, 0x12, 0x34, 0x56, 0x78, 0x90, 0x01, 0x12, 0x34, 0x56, 0x78, 0x90, 0x02, 0x12, 0x34, 0x56, 0x78, 0x90, 0x03, 0x12, 0x34, 0x56, 0x78, 0x90, 0x04,
				0x03, 0x07, 0xf9, 0xff,
				0x03, 0x07, 0xf9, 0xff,
				0x00, 0x91, 0x00, 0x21,
				0x04, 0x12, 0x34, 0x56, 0x78, 0x90, 0x01, 0x12, 0x34, 0x56, 0x78, 0x90, 0x02, 0x12, 0x34, 0x56, 0x78, 0x90, 0x03, 0x12, 0x34, 0x56, 0x78, 0x90, 0x04,
				0x03, 0x07, 0xf9, 0xff,
				0x03, 0x07, 0xf9, 0xff,
				0x00, 0x7e, 0x00, 0x02, 0x80, 0xff,
			},
		}, {
			Description: "Multiple IEs",
			Structured: message.NewSessionDeletionResponse(
				mp, fo, seid, seq, pri,
				ie.NewCause(ie.CauseRequestAccepted),
				ie.NewOffendingIE(ie.Cause),
				ie.NewLoadControlInformation(ie.NewSequenceNumber(0xffffffff), ie.NewMetric(0x01)),
				ie.NewOverloadControlInformation(
					ie.NewSequenceNumber(0xffffffff),
					ie.NewMetric(0x01),
					ie.NewTimer(20*time.Hour),
					ie.NewOCIFlags(0x01),
				),
				ie.NewUsageReportWithinSessionDeletionResponse(
					ie.NewURRID(0xffffffff),
					ie.NewURSEQN(0xffffffff),
					ie.NewUsageReportTrigger(0xff, 0xff, 0xff),
					ie.NewStartTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewEndTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewVolumeMeasurement(0x3f, 0x1111111111111111, 0x2222222222222222, 0x3333333333333333, 0x4444444444444444, 0x5555555555555555, 0x6666666666666666),
					ie.NewDurationMeasurement(10*time.Second),
					ie.NewTimeOfFirstPacket(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewTimeOfLastPacket(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewUsageInformation(1, 1, 1, 1),
					ie.NewEthernetTrafficInformation(
						ie.NewMACAddressesDetected(
							ie.NewCTAG(0x07, 1, 1, 4095),
							ie.NewSTAG(0x07, 1, 1, 4095),
							mac1, mac2, mac3, mac4,
						),
						ie.NewMACAddressesRemoved(
							ie.NewCTAG(0x07, 1, 1, 4095),
							ie.NewSTAG(0x07, 1, 1, 4095),
							mac1, mac2, mac3, mac4,
						),
					),
				),
				ie.NewUsageReportWithinSessionDeletionResponse(
					ie.NewURRID(0xeeeeeeee),
					ie.NewURSEQN(0xffffffff),
					ie.NewUsageReportTrigger(0xff, 0xff, 0xff),
					ie.NewStartTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewEndTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewVolumeMeasurement(0x3f, 0x1111111111111111, 0x2222222222222222, 0x3333333333333333, 0x4444444444444444, 0x5555555555555555, 0x6666666666666666),
					ie.NewDurationMeasurement(10*time.Second),
					ie.NewTimeOfFirstPacket(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewTimeOfLastPacket(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewUsageInformation(1, 1, 1, 1),
					ie.NewEthernetTrafficInformation(
						ie.NewMACAddressesDetected(
							ie.NewCTAG(0x07, 1, 1, 4095),
							ie.NewSTAG(0x07, 1, 1, 4095),
							mac1, mac2, mac3, mac4,
						),
						ie.NewMACAddressesRemoved(
							ie.NewCTAG(0x07, 1, 1, 4095),
							ie.NewSTAG(0x07, 1, 1, 4095),
							mac1, mac2, mac3, mac4,
						),
					),
				),
				ie.NewAdditionalUsageReportsInformation(0x00ff),
			),
			Serialized: []byte{
				0x21, 0x37, 0x01, 0xdf, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x11, 0x22, 0x33, 0x00,
				0x00, 0x13, 0x00, 0x01, 0x01,
				0x00, 0x28, 0x00, 0x02, 0x00, 0x13,
				0x00, 0x33, 0x00, 0x0d, 0x00, 0x34, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff, 0x00, 0x35, 0x00, 0x01, 0x01,
				0x00, 0x36, 0x00, 0x17,
				0x00, 0x34, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0x35, 0x00, 0x01, 0x01,
				0x00, 0x37, 0x00, 0x01, 0x82,
				0x00, 0x6e, 0x00, 0x01, 0x01,
				0x00, 0x4f, 0x00, 0xc7,
				0x00, 0x51, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0x68, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0x3f, 0x00, 0x03, 0xff, 0xff, 0xff,
				0x00, 0x4b, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x4c, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x42, 0x00, 0x31,
				0x3f,
				0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11,
				0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22,
				0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33,
				0x44, 0x44, 0x44, 0x44, 0x44, 0x44, 0x44, 0x44,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
				0x00, 0x43, 0x00, 0x04, 0x00, 0x00, 0x00, 0x0a,
				0x00, 0x45, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x46, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x5a, 0x00, 0x01, 0x0f,
				0x00, 0x8f, 0x00, 0x4a,
				0x00, 0x90, 0x00, 0x21,
				0x04, 0x12, 0x34, 0x56, 0x78, 0x90, 0x01, 0x12, 0x34, 0x56, 0x78, 0x90, 0x02, 0x12, 0x34, 0x56, 0x78, 0x90, 0x03, 0x12, 0x34, 0x56, 0x78, 0x90, 0x04,
				0x03, 0x07, 0xf9, 0xff,
				0x03, 0x07, 0xf9, 0xff,
				0x00, 0x91, 0x00, 0x21,
				0x04, 0x12, 0x34, 0x56, 0x78, 0x90, 0x01, 0x12, 0x34, 0x56, 0x78, 0x90, 0x02, 0x12, 0x34, 0x56, 0x78, 0x90, 0x03, 0x12, 0x34, 0x56, 0x78, 0x90, 0x04,
				0x03, 0x07, 0xf9, 0xff,
				0x03, 0x07, 0xf9, 0xff,
				0x00, 0x4f, 0x00, 0xc7,
				0x00, 0x51, 0x00, 0x04, 0xee, 0xee, 0xee, 0xee,
				0x00, 0x68, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff,
				0x00, 0x3f, 0x00, 0x03, 0xff, 0xff, 0xff,
				0x00, 0x4b, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x4c, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x42, 0x00, 0x31,
				0x3f,
				0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11,
				0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22,
				0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33,
				0x44, 0x44, 0x44, 0x44, 0x44, 0x44, 0x44, 0x44,
				0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55,
				0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
				0x00, 0x43, 0x00, 0x04, 0x00, 0x00, 0x00, 0x0a,
				0x00, 0x45, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x46, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00,
				0x00, 0x5a, 0x00, 0x01, 0x0f,
				0x00, 0x8f, 0x00, 0x4a,
				0x00, 0x90, 0x00, 0x21,
				0x04, 0x12, 0x34, 0x56, 0x78, 0x90, 0x01, 0x12, 0x34, 0x56, 0x78, 0x90, 0x02, 0x12, 0x34, 0x56, 0x78, 0x90, 0x03, 0x12, 0x34, 0x56, 0x78, 0x90, 0x04,
				0x03, 0x07, 0xf9, 0xff,
				0x03, 0x07, 0xf9, 0xff,
				0x00, 0x91, 0x00, 0x21,
				0x04, 0x12, 0x34, 0x56, 0x78, 0x90, 0x01, 0x12, 0x34, 0x56, 0x78, 0x90, 0x02, 0x12, 0x34, 0x56, 0x78, 0x90, 0x03, 0x12, 0x34, 0x56, 0x78, 0x90, 0x04,
				0x03, 0x07, 0xf9, 0xff,
				0x03, 0x07, 0xf9, 0xff,
				0x00, 0x7e, 0x00, 0x02, 0x80, 0xff,
			},
		},
	}

	testutil.Run(t, cases, func(b []byte) (testutil.Serializable, error) {
		v, err := message.ParseSessionDeletionResponse(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
