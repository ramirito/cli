package resources

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func cf_i18n_test_fixtures_main_en_us_all_json() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x8a, 0xae,
		0x56, 0xca, 0x4c, 0x51, 0xb2, 0x52, 0xf2, 0x48, 0xcd, 0xc9, 0xc9, 0x57,
		0x28, 0xcf, 0x2f, 0xca, 0x49, 0x51, 0x54, 0xd2, 0x51, 0x2a, 0x29, 0x4a,
		0xcc, 0x2b, 0xce, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x43, 0x97, 0xac, 0xd5,
		0x81, 0x6a, 0xf1, 0xf1, 0xf7, 0xc1, 0x50, 0xe9, 0x93, 0x58, 0x9a, 0x9e,
		0xa1, 0x90, 0x5f, 0x5a, 0xa2, 0x90, 0x93, 0x5f, 0x9a, 0xa2, 0x54, 0x1b,
		0xcb, 0x05, 0x08, 0x00, 0x00, 0xff, 0xff, 0x95, 0x67, 0x11, 0xae, 0x61,
		0x00, 0x00, 0x00,
	},
		"cf/i18n/test_fixtures/main/en_US.all.json",
	)
}

func cf_i18n_test_fixtures_main_fr_fr_all_json() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x8a, 0xae,
		0x56, 0xca, 0x4c, 0x51, 0xb2, 0x52, 0xf2, 0x48, 0xcd, 0xc9, 0xc9, 0x57,
		0x28, 0xcf, 0x2f, 0xca, 0x49, 0x51, 0x54, 0xd2, 0x51, 0x2a, 0x29, 0x4a,
		0xcc, 0x2b, 0xce, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x03, 0x4a, 0x1e, 0x6e,
		0x00, 0xca, 0xe5, 0xa4, 0x2a, 0xe4, 0xe6, 0xe7, 0xa5, 0xa4, 0x2a, 0x2a,
		0xd5, 0xc6, 0x72, 0x01, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x0b, 0x22,
		0xb2, 0x37, 0x00, 0x00, 0x00,
	},
		"cf/i18n/test_fixtures/main/fr_FR.all.json",
	)
}

func cf_i18n_test_fixtures_main_hello_go() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x2a, 0x48,
		0x4c, 0xce, 0x4e, 0x4c, 0x4f, 0x55, 0xc8, 0x4d, 0xcc, 0xcc, 0xe3, 0xe2,
		0xca, 0xcc, 0x2d, 0xc8, 0x2f, 0x2a, 0x51, 0x50, 0x4a, 0xcb, 0x2d, 0x51,
		0xe2, 0xe2, 0x4a, 0x2b, 0xcd, 0x4b, 0x06, 0x4b, 0x68, 0x68, 0x2a, 0x54,
		0x73, 0x71, 0x02, 0x05, 0xf5, 0x02, 0x8a, 0x32, 0xf3, 0x4a, 0x72, 0xf2,
		0x34, 0x94, 0x3c, 0x52, 0x73, 0x72, 0xf2, 0x15, 0xca, 0xf3, 0x8b, 0x72,
		0x52, 0x14, 0x95, 0x34, 0xd1, 0x24, 0x7d, 0xfc, 0x7d, 0x80, 0x62, 0xb5,
		0x5c, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x13, 0x28, 0xd3, 0xfe, 0x5d,
		0x00, 0x00, 0x00,
	},
		"cf/i18n/test_fixtures/main/hello.go",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	if f, ok := _bindata[name]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"cf/i18n/test_fixtures/main/en_US.all.json": cf_i18n_test_fixtures_main_en_us_all_json,
	"cf/i18n/test_fixtures/main/fr_FR.all.json": cf_i18n_test_fixtures_main_fr_fr_all_json,
	"cf/i18n/test_fixtures/main/hello.go":       cf_i18n_test_fixtures_main_hello_go,
}