// SPDX-License-Identifier: MIT

package generated_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/GDATAAdvancedAnalytics/winreg-tasks/golang/generated"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
)

func TestTaskParser_ExeTasks(t *testing.T) {
	blob := []byte{
		0x03, 0x00,
		0x0c, 0x00, 0x00, 0x00,
		0x41, 0x00, 0x75, 0x00, 0x74, 0x00, 0x68, 0x00, 0x6f, 0x00, 0x72, 0x00,

		0x66, 0x66,
		0x08, 0x00, 0x00, 0x00, 0x65, 0x00, 0x78, 0x00, 0x65, 0x00, 0x63, 0x00,
		0x1a, 0x00, 0x00, 0x00, 0x65, 0x00, 0x78, 0x00, 0x65, 0x00, 0x63, 0x00, 0x75, 0x00, 0x74, 0x00, 0x69, 0x00, 0x6f, 0x00, 0x6e, 0x00, 0x2e, 0x00, 0x65, 0x00, 0x78, 0x00, 0x65, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x01, 0x00,

		0x66, 0x66,
		0x08, 0x00, 0x00, 0x00, 0x65, 0x00, 0x78, 0x00, 0x65, 0x00, 0x63, 0x00,
		0x10, 0x00, 0x00, 0x00, 0x6d, 0x00, 0x6f, 0x00, 0x61, 0x00, 0x72, 0x00, 0x2e, 0x00, 0x65, 0x00, 0x78, 0x00, 0x65, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00,

		0x66, 0x66,
		0x08, 0x00, 0x00, 0x00, 0x65, 0x00, 0x78, 0x00, 0x65, 0x00, 0x63, 0x00,
		0x42, 0x00, 0x00, 0x00, 0x43, 0x00, 0x3a, 0x00, 0x5c, 0x00, 0x75, 0x00, 0x73, 0x00, 0x65, 0x00, 0x72, 0x00, 0x73, 0x00, 0x5c, 0x00, 0x63, 0x00, 0x61, 0x00, 0x73, 0x00, 0x65, 0x00, 0x5c, 0x00, 0x44, 0x00, 0x65, 0x00, 0x73, 0x00, 0x6b, 0x00, 0x74, 0x00, 0x6f, 0x00, 0x70, 0x00, 0x5c, 0x00, 0x77, 0x00, 0x69, 0x00, 0x74, 0x00, 0x68, 0x00, 0x61, 0x00, 0x72, 0x00, 0x67, 0x00, 0x2e, 0x00, 0x78, 0x00, 0x6d, 0x00, 0x6c, 0x00,
		0x10, 0x00, 0x00, 0x00, 0x61, 0x00, 0x72, 0x00, 0x67, 0x00, 0x20, 0x00, 0x61, 0x00, 0x72, 0x00, 0x67, 0x00, 0x32, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00,

		0x66, 0x66,
		0x08, 0x00, 0x00, 0x00, 0x65, 0x00, 0x78, 0x00, 0x65, 0x00, 0x63, 0x00,
		0x2c, 0x00, 0x00, 0x00, 0x44, 0x00, 0x3a, 0x00, 0x5c, 0x00, 0x6f, 0x00, 0x6e, 0x00, 0x6c, 0x00, 0x79, 0x00, 0x5f, 0x00, 0x77, 0x00, 0x6f, 0x00, 0x72, 0x00, 0x6b, 0x00, 0x69, 0x00, 0x6e, 0x00, 0x67, 0x00, 0x64, 0x00, 0x69, 0x00, 0x72, 0x00, 0x2e, 0x00, 0x70, 0x00, 0x73, 0x00, 0x31, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x14, 0x00, 0x00, 0x00, 0x43, 0x00, 0x3a, 0x00, 0x5c, 0x00, 0x77, 0x00, 0x69, 0x00, 0x6e, 0x00, 0x64, 0x00, 0x6f, 0x00, 0x77, 0x00, 0x73, 0x00,
		0x00, 0x00,

		0x66, 0x66,
		0x08, 0x00, 0x00, 0x00, 0x65, 0x00, 0x78, 0x00, 0x65, 0x00, 0x63, 0x00,
		0x1e, 0x00, 0x00, 0x00, 0x7a, 0x00, 0x3a, 0x00, 0x5c, 0x00, 0x61, 0x00, 0x6c, 0x00, 0x6c, 0x00, 0x6f, 0x00, 0x66, 0x00, 0x5c, 0x00, 0x69, 0x00, 0x74, 0x00, 0x2e, 0x00, 0x65, 0x00, 0x78, 0x00, 0x65, 0x00,
		0x12, 0x00, 0x00, 0x00, 0x61, 0x00, 0x72, 0x00, 0x67, 0x00, 0x31, 0x00, 0x20, 0x00, 0x61, 0x00, 0x72, 0x00, 0x67, 0x00, 0x32, 0x00,
		0x1a, 0x00, 0x00, 0x00, 0x63, 0x00, 0x3a, 0x00, 0x5c, 0x00, 0x77, 0x00, 0x6f, 0x00, 0x72, 0x00, 0x6b, 0x00, 0x69, 0x00,
		0x6e, 0x00, 0x67, 0x00, 0x64, 0x00, 0x69, 0x00, 0x72,
		0x00, 0x00,
	}

	expected := []*generated.Actions_ExeTaskProperties{
		{Command: &generated.Bstr{Str: "execution.exe"}, Arguments: &generated.Bstr{Str: ""}, WorkingDirectory: &generated.Bstr{Str: ""}, Flags: uint16(1)},
		{Command: &generated.Bstr{Str: "moar.exe"}, Arguments: &generated.Bstr{Str: ""}, WorkingDirectory: &generated.Bstr{Str: ""}, Flags: uint16(0)},
		{Command: &generated.Bstr{Str: `C:\users\case\Desktop\witharg.xml`}, Arguments: &generated.Bstr{Str: "arg arg2"}, WorkingDirectory: &generated.Bstr{Str: ""}, Flags: uint16(0)},
		{Command: &generated.Bstr{Str: `D:\only_workingdir.ps1`}, Arguments: &generated.Bstr{Str: ""}, WorkingDirectory: &generated.Bstr{Str: `C:\windows`}, Flags: uint16(0)},
		{Command: &generated.Bstr{Str: `z:\allof\it.exe`}, Arguments: &generated.Bstr{Str: "arg1 arg2"}, WorkingDirectory: &generated.Bstr{Str: `c:\workingdir`}, Flags: uint16(0)},
	}

	tt := generated.NewActions()
	if err := tt.Read(kaitai.NewStream(bytes.NewReader(blob)), tt, tt); err != nil {
		t.Error(fmt.Sprintf("Read (%v)", err))
	}

	if tt.Context.Str != "Author" {
		t.Error("Context")
	}

	if len(tt.Actions) != 5 {
		t.Error("Amount of Tasks")
	}

	for i, task := range tt.Actions {
		e := &generated.Actions_Action{
			Id:         &generated.Bstr{Str: "exec"},
			Magic:      uint16(0x6666),
			Properties: expected[i],
		}
		if err := compareExeTask(task, e); err != nil {
			t.Errorf("Task %d (%v)", i, err)
		}
	}
}

func TestTaskParser_ComHandler(t *testing.T) {
	blob := []byte{
		0x03, 0x00,
		0x0a, 0x00, 0x00, 0x00, 0x55, 0x00, 0x73, 0x00, 0x65, 0x00, 0x72, 0x00, 0x73, 0x00,

		0x77, 0x77,
		0x14, 0x00, 0x00, 0x00, 0x63, 0x00, 0x6f, 0x00, 0x6d, 0x00, 0x68, 0x00, 0x61, 0x00, 0x6e, 0x00, 0x64, 0x00, 0x6c, 0x00, 0x65, 0x00, 0x72, 0x00,
		0xb9, 0xd1, 0xbc, 0x61, 0x0c, 0x34, 0xec, 0x40, 0x9d, 0x41, 0xd7, 0xf1, 0xc0, 0x63, 0x2f, 0x05,
		0x24, 0x00, 0x00, 0x00, 0x4d, 0x00, 0x69, 0x00, 0x73, 0x00, 0x73, 0x00, 0x69, 0x00, 0x6e, 0x00, 0x67, 0x00, 0x43, 0x00, 0x72, 0x00, 0x65, 0x00, 0x64, 0x00, 0x65, 0x00, 0x6e, 0x00, 0x74, 0x00, 0x69, 0x00, 0x61, 0x00, 0x6c, 0x00, 0x73, 0x00,
	}

	expected_task := &generated.Actions_Action{
		Id:    &generated.Bstr{Str: "comhandler"},
		Magic: uint16(0x7777),
		Properties: &generated.Actions_ComHandlerProperties{
			Clsid: []byte{0xb9, 0xd1, 0xbc, 0x61, 0x0c, 0x34, 0xec, 0x40, 0x9d, 0x41, 0xd7, 0xf1, 0xc0, 0x63, 0x2f, 0x05},
			Data:  &generated.Bstr{Str: "MissingCredentials"},
		},
	}

	tt := generated.NewActions()
	if err := tt.Read(kaitai.NewStream(bytes.NewReader(blob)), tt, tt); err != nil {
		t.Error(fmt.Sprintf("Read (%v)", err))
		t.FailNow()
	}

	if tt.Context.Str != "Users" {
		t.Error("Context")
	}

	if len(tt.Actions) != 1 {
		t.Error("Amount of Tasks")
	}

	if err := compareComHandlerTask(tt.Actions[0], expected_task); err != nil {
		t.Error(err)
	}
}

func TestTaskParser_MessageBox(t *testing.T) {
	blob := []byte{
		0x03, 0x00,
		0x0c, 0x00, 0x00, 0x00, 0x41, 0x00, 0x75, 0x00, 0x74, 0x00, 0x68, 0x00, 0x6f, 0x00, 0x72, 0x00,

		0x99, 0x99,
		0x14, 0x00, 0x00, 0x00, 0x6d, 0x00, 0x65, 0x00, 0x73, 0x00, 0x73, 0x00, 0x61, 0x00, 0x67, 0x00, 0x65, 0x00, 0x62, 0x00, 0x6f, 0x00, 0x78, 0x00,
		0x1e, 0x00, 0x00, 0x00, 0x7a, 0x00, 0x3a, 0x00, 0x5c, 0x00, 0x61, 0x00, 0x6c, 0x00, 0x6c, 0x00, 0x6f, 0x00, 0x66, 0x00, 0x5c, 0x00, 0x69, 0x00, 0x74, 0x00, 0x2e, 0x00, 0x65, 0x00, 0x78, 0x00, 0x65, 0x00,
		0x12, 0x00, 0x00, 0x00, 0x61, 0x00, 0x72, 0x00, 0x67, 0x00, 0x31, 0x00, 0x20, 0x00, 0x61, 0x00, 0x72, 0x00, 0x67, 0x00, 0x32, 0x00,
	}

	expected_task := &generated.Actions_Action{
		Id:    &generated.Bstr{Str: "messagebox"},
		Magic: uint16(0x9999),
		Properties: &generated.Actions_MessageboxTaskProperties{
			Caption: &generated.Bstr{Str: `z:\allof\it.exe`},
			Content: &generated.Bstr{Str: "arg1 arg2"},
		},
	}

	tt := generated.NewActions()
	if err := tt.Read(kaitai.NewStream(bytes.NewReader(blob)), tt, tt); err != nil {
		t.Error(fmt.Sprintf("Read (%v)", err))
		t.FailNow()
	}

	if tt.Context.Str != "Author" {
		t.Error("Context")
	}

	if len(tt.Actions) != 1 {
		t.Error("Amount of Tasks")
	}

	if err := compareMessageBoxTask(tt.Actions[0], expected_task); err != nil {
		t.Error(err)
	}
}

func TestTaskParser_EmailTask(t *testing.T) {
	blob := []byte{
		0x03, 0x00,
		0x0c, 0x00, 0x00, 0x00, 0x41, 0x00, 0x75, 0x00, 0x74, 0x00, 0x68, 0x00, 0x6f, 0x00, 0x72, 0x00,

		0x88, 0x88,
		0x0a, 0x00, 0x00, 0x00, 0x65, 0x00, 0x6d, 0x00, 0x61, 0x00, 0x69, 0x00, 0x6c, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x62, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x63, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x64, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x65, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x66, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x67, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x68, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x69, 0x00,
		0x03, 0x00, 0x00, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x6a, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x6b, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x6c, 0x00,
		0x02, 0x00, 0x00, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x6d, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x6e, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x6f, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x70, 0x00,
	}

	expected_task := &generated.Actions_Action{
		Id:    &generated.Bstr{Str: "email"},
		Magic: uint16(0x8888),
		Properties: &generated.Actions_EmailTaskProperties{
			Server:  &generated.Bstr{Str: "g"},
			Subject: &generated.Bstr{Str: "h"},
			To:      &generated.Bstr{Str: "c"},
			Cc:      &generated.Bstr{Str: "d"},
			Bcc:     &generated.Bstr{Str: "e"},
			ReplyTo: &generated.Bstr{Str: "f"},
			From:    &generated.Bstr{Str: "b"},
			Body:    &generated.Bstr{Str: "i"},
			Headers: []*generated.Actions_KeyValueString{
				{Key: &generated.Bstr{Str: "m"}, Value: &generated.Bstr{Str: "n"}},
				{Key: &generated.Bstr{Str: "o"}, Value: &generated.Bstr{Str: "p"}},
			},
			AttachmentFilenames: []*generated.Bstr{
				{Str: "j"},
				{Str: "k"},
				{Str: "l"},
			},
		},
	}

	tt := generated.NewActions()
	if err := tt.Read(kaitai.NewStream(bytes.NewReader(blob)), tt, tt); err != nil {
		t.Error(fmt.Sprintf("Read (%v)", err))
		t.FailNow()
	}

	if tt.Context.Str != "Author" {
		t.Error("Context")
	}

	if len(tt.Actions) != 1 {
		t.Error("Amount of Tasks")
	}

	if err := compareEmailTask(tt.Actions[0], expected_task); err != nil {
		t.Error(err)
	}
}

func TestTaskParser_MixedAll(t *testing.T) {
	blob := []byte{
		0x03, 0x00,
		0x0c, 0x00, 0x00, 0x00, 0x41, 0x00, 0x75, 0x00, 0x74, 0x00, 0x68, 0x00, 0x6f, 0x00, 0x72, 0x00,

		0x66, 0x66,
		0x08, 0x00, 0x00, 0x00, 0x65, 0x00, 0x78, 0x00, 0x65, 0x00, 0x63, 0x00,
		0x1a, 0x00, 0x00, 0x00, 0x65, 0x00, 0x78, 0x00, 0x65, 0x00, 0x63, 0x00, 0x75, 0x00, 0x74, 0x00, 0x69, 0x00, 0x6f, 0x00, 0x6e, 0x00, 0x2e, 0x00, 0x65, 0x00, 0x78, 0x00, 0x65, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x01, 0x00,

		0x77, 0x77,
		0x14, 0x00, 0x00, 0x00, 0x63, 0x00, 0x6f, 0x00, 0x6d, 0x00, 0x68, 0x00, 0x61, 0x00, 0x6e, 0x00, 0x64, 0x00, 0x6c, 0x00, 0x65, 0x00, 0x72, 0x00,
		0xb9, 0xd1, 0xbc, 0x61, 0x0c, 0x34, 0xec, 0x40, 0x9d, 0x41, 0xd7, 0xf1, 0xc0, 0x63, 0x2f, 0x05,
		0x24, 0x00, 0x00, 0x00, 0x4d, 0x00, 0x69, 0x00, 0x73, 0x00, 0x73, 0x00, 0x69, 0x00, 0x6e, 0x00, 0x67, 0x00, 0x43, 0x00, 0x72, 0x00, 0x65, 0x00, 0x64, 0x00, 0x65, 0x00, 0x6e, 0x00, 0x74, 0x00, 0x69, 0x00, 0x61, 0x00, 0x6c, 0x00, 0x73, 0x00,

		0x88, 0x88,
		0x0a, 0x00, 0x00, 0x00, 0x65, 0x00, 0x6d, 0x00, 0x61, 0x00, 0x69, 0x00, 0x6c, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x62, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x63, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x64, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x65, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x66, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x67, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x68, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x69, 0x00,
		0x03, 0x00, 0x00, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x6a, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x6b, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x6c, 0x00,
		0x02, 0x00, 0x00, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x6d, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x6e, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x6f, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x70, 0x00,

		0x99, 0x99,
		0x14, 0x00, 0x00, 0x00, 0x6d, 0x00, 0x65, 0x00, 0x73, 0x00, 0x73, 0x00, 0x61, 0x00, 0x67, 0x00, 0x65, 0x00, 0x62, 0x00, 0x6f, 0x00, 0x78, 0x00,
		0x1e, 0x00, 0x00, 0x00, 0x7a, 0x00, 0x3a, 0x00, 0x5c, 0x00, 0x61, 0x00, 0x6c, 0x00, 0x6c, 0x00, 0x6f, 0x00, 0x66, 0x00, 0x5c, 0x00, 0x69, 0x00, 0x74, 0x00, 0x2e, 0x00, 0x65, 0x00, 0x78, 0x00, 0x65, 0x00,
		0x12, 0x00, 0x00, 0x00, 0x61, 0x00, 0x72, 0x00, 0x67, 0x00, 0x31, 0x00, 0x20, 0x00, 0x61, 0x00, 0x72, 0x00, 0x67, 0x00, 0x32, 0x00,
	}

	expected_exe := &generated.Actions_Action{
		Id:    &generated.Bstr{Str: "exec"},
		Magic: uint16(0x6666),
		Properties: &generated.Actions_ExeTaskProperties{
			Command:          &generated.Bstr{Str: "execution.exe"},
			Arguments:        &generated.Bstr{Str: ""},
			WorkingDirectory: &generated.Bstr{Str: ""},
			Flags:            uint16(1),
		},
	}

	expected_comhandler := &generated.Actions_Action{
		Id:    &generated.Bstr{Str: "comhandler"},
		Magic: uint16(0x7777),
		Properties: &generated.Actions_ComHandlerProperties{
			Clsid: []byte{0xb9, 0xd1, 0xbc, 0x61, 0x0c, 0x34, 0xec, 0x40, 0x9d, 0x41, 0xd7, 0xf1, 0xc0, 0x63, 0x2f, 0x05},
			Data:  &generated.Bstr{Str: "MissingCredentials"},
		},
	}

	expected_email := &generated.Actions_Action{
		Id:    &generated.Bstr{Str: "email"},
		Magic: uint16(0x8888),
		Properties: &generated.Actions_EmailTaskProperties{
			Server:  &generated.Bstr{Str: "g"},
			Subject: &generated.Bstr{Str: "h"},
			To:      &generated.Bstr{Str: "c"},
			Cc:      &generated.Bstr{Str: "d"},
			Bcc:     &generated.Bstr{Str: "e"},
			ReplyTo: &generated.Bstr{Str: "f"},
			From:    &generated.Bstr{Str: "b"},
			Body:    &generated.Bstr{Str: "i"},
			Headers: []*generated.Actions_KeyValueString{
				{Key: &generated.Bstr{Str: "m"}, Value: &generated.Bstr{Str: "n"}},
				{Key: &generated.Bstr{Str: "o"}, Value: &generated.Bstr{Str: "p"}},
			},
			AttachmentFilenames: []*generated.Bstr{
				{Str: "j"},
				{Str: "k"},
				{Str: "l"},
			},
		},
	}

	expected_messagebox := &generated.Actions_Action{
		Id:    &generated.Bstr{Str: "messagebox"},
		Magic: uint16(0x9999),
		Properties: &generated.Actions_MessageboxTaskProperties{
			Caption: &generated.Bstr{Str: `z:\allof\it.exe`},
			Content: &generated.Bstr{Str: "arg1 arg2"},
		},
	}

	tt := generated.NewActions()
	if err := tt.Read(kaitai.NewStream(bytes.NewReader(blob)), tt, tt); err != nil {
		t.Error(fmt.Sprintf("Read (%v)", err))
		t.FailNow()
	}

	if tt.Context.Str != "Author" {
		t.Error("Context")
	}

	if len(tt.Actions) != 4 {
		t.Error("Amount of Tasks")
	}

	if err := compareExeTask(tt.Actions[0], expected_exe); err != nil {
		t.Errorf("ExecTask failed: %v", err)
	}

	if err := compareComHandlerTask(tt.Actions[1], expected_comhandler); err != nil {
		t.Errorf("ComHandler failed: %v", err)
	}

	if err := compareEmailTask(tt.Actions[2], expected_email); err != nil {
		t.Errorf("EmailTask failed: %v", err)
	}

	if err := compareMessageBoxTask(tt.Actions[3], expected_messagebox); err != nil {
		t.Errorf("MesageBoxTask failed: %v", err)
	}
}
