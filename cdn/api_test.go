package cdn

import (
	"github.com/qiniu/api.v7/auth/qbox"
	"os"
	"testing"
	"time"
)

//global variables

var (
	ak     = os.Getenv("QINIU_ACCESS_KEY")
	sk     = os.Getenv("QINIU_SECRET_KEY")
	domain = os.Getenv("QINIU_TEST_DOMAIN")

	layout    = "2006-01-02"
	now       = time.Now()
	startDate = now.AddDate(0, 0, -2).Format(layout)
	endDate   = now.AddDate(0, 0, -1).Format(layout)
	logDate   = now.AddDate(0, 0, -1).Format(layout)

	testUrls = []string{
		"http://qdisk.qiniudn.com/qiniu1.png",
		"http://qdisk.qiniudn.com/qiniu2.png",
	}
	testDirs = []string{
		"http://qdisk.qiniudn.com/dir1/",
		"http://qdisk.qiniudn.com/dir2/",
	}
)

var mac *qbox.Mac
var cdnManager *CdnManager

func init() {
	if ak == "" || sk == "" {
		panic("please run ./test-env.sh first")
	}
	mac = qbox.NewMac(ak, sk)
	cdnManager = &CdnManager{Mac: mac}
}

//TestGetBandwidthData
func TestGetBandwidthData(t *testing.T) {
	type args struct {
		startDate   string
		endDate     string
		granularity string
		domainList  []string
	}

	testCases := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "CdnManager_TestGetBandwidthData",
			args: args{
				startDate,
				endDate,
				"5min",
				[]string{domain},
			},
			wantCode: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ret, err := cdnManager.GetBandwidthData(tc.args.startDate, tc.args.endDate,
				tc.args.granularity, tc.args.domainList)
			if err != nil || ret.Code != tc.wantCode {
				t.Errorf("GetBandwidth() error = %v, %v", err, ret.Error)
				return
			}
		})
	}
}

//TestGetFluxData
func TestGetFluxData(t *testing.T) {
	type args struct {
		startDate   string
		endDate     string
		granularity string
		domainList  []string
	}

	testCases := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "CdnManager_TestGetFluxData",
			args: args{
				startDate,
				endDate,
				"5min",
				[]string{domain},
			},
			wantCode: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ret, err := cdnManager.GetFluxData(tc.args.startDate, tc.args.endDate,
				tc.args.granularity, tc.args.domainList)
			if err != nil || ret.Code != tc.wantCode {
				t.Errorf("GetFlux() error = %v, %v", err, ret.Error)
				return
			}
		})
	}
}

//TestRefreshUrls
func TestRefreshUrls(t *testing.T) {
	type args struct {
		urls []string
	}

	testCases := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "CdnManager_TestRefresUrls",
			args: args{
				urls: testUrls,
			},
			wantCode: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ret, err := cdnManager.RefreshUrls(tc.args.urls)
			if err != nil || ret.Code != tc.wantCode {
				t.Errorf("RefreshUrls() error = %v, %v", err, ret.Error)
				return
			}
		})
	}
}

//TestRefreshDirs
func TestRefreshDirs(t *testing.T) {
	type args struct {
		dirs []string
	}

	testCases := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "CdnManager_TestRefreshDirs",
			args: args{
				dirs: testDirs,
			},
			wantCode: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ret, err := cdnManager.RefreshDirs(tc.args.dirs)
			if err != nil || ret.Code != tc.wantCode {
				t.Errorf("RefreshDirs() error = %v, %v", err, ret.Error)
				return
			}
		})
	}
}

//TestPrefetchUrls
func TestPrefetchUrls(t *testing.T) {
	type args struct {
		urls []string
	}

	testCases := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "CdnManager_PrefetchUrls",
			args: args{
				urls: testUrls,
			},
			wantCode: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ret, err := cdnManager.PrefetchUrls(tc.args.urls)
			if err != nil || ret.Code != tc.wantCode {
				t.Errorf("PrefetchUrls() error = %v, %v", err, ret.Error)
				return
			}
		})
	}
}

//TestGetCdnLogList
func TestGetCdnLogList(t *testing.T) {
	type args struct {
		date    string
		domains string
	}

	testCases := []struct {
		name string
		args args
	}{
		{
			name: "CdnManager_TestGetCdnLogList",
			args: args{
				date:    logDate,
				domains: domain,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := cdnManager.GetCdnLogList(tc.args.date, tc.args.domains)
			if err != nil {
				t.Errorf("GetCdnLogList() error = %v, %v", err)
				return
			}
		})
	}
}
