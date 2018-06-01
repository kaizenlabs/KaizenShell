// Complete list of command line flags for versioninfo.json generation:
//   -charset=0: charset ID
//   -comment="": StringFileInfo.Comments
//   -company="": StringFileInfo.CompanyName
//   -copyright="": StringFileInfo.LegalCopyright
//   -description="": StringFileInfo.FileDescription
//   -example=false: just dump out an example versioninfo.json to stdout
//   -file-version="": StringFileInfo.FileVersion
//   -icon="": icon file name
//   -internal-name="": StringFileInfo.InternalName
//   -manifest="": manifest file name
//   -o="resource.syso": output file name
//   -platform-specific=false: output i386 and amd64 named resource.syso, ignores -o
//   -original-name="": StringFileInfo.OriginalFilename
//   -private-build="": StringFileInfo.PrivateBuild
//   -product-name="": StringFileInfo.ProductName
//   -product-version="": StringFileInfo.ProductVersion
//   -special-build="": StringFileInfo.SpecialBuild
//   -trademark="": StringFileInfo.LegalTrademarks
//   -translation=0: translation ID
//   -64:false: generate 64-bit binaries on true
//   -ver-major=-1: FileVersion.Major
//   -ver-minor=-1: FileVersion.Minor
//   -ver-patch=-1: FileVersion.Patch
//   -ver-build=-1: FileVersion.Build
//   -product-ver-major=-1: ProductVersion.Major
//   -product-ver-minor=-1: ProductVersion.Minor
//   -product-ver-patch=-1: ProductVersion.Patch
//   -product-ver-build=-1: ProductVersion.Build

package main

import (
	"flag"
)

// VersionInfo is a struct representing the final output file structure in JSON format
type VersionInfo struct {
	FixedFileInfo struct {
		FileVersion struct {
			Major int `json:"Major"`
			Minor int `json:"Minor"`
			Patch int `json:"Patch"`
			Build int `json:"Build"`
		} `json:"FileVersion"`
		ProductVersion struct {
			Major int `json:"Major"`
			Minor int `json:"Minor"`
			Patch int `json:"Patch"`
			Build int `json:"Build"`
		} `json:"ProductVersion"`
		FileFlagsMask string `json:"FileFlagsMask"`
		FileFlags     string `json:"FileFlags "`
		FileOS        string `json:"FileOS"`
		FileType      string `json:"FileType"`
		FileSubType   string `json:"FileSubType"`
	} `json:"FixedFileInfo"`
	StringFileInfo struct {
		Comments         string `json:"Comments"`
		CompanyName      string `json:"CompanyName"`
		FileDescription  string `json:"FileDescription"`
		FileVersion      string `json:"FileVersion"`
		InternalName     string `json:"InternalName"`
		LegalCopyright   string `json:"LegalCopyright"`
		LegalTrademarks  string `json:"LegalTrademarks"`
		OriginalFilename string `json:"OriginalFilename"`
		PrivateBuild     string `json:"PrivateBuild"`
		ProductName      string `json:"ProductName"`
		ProductVersion   string `json:"ProductVersion"`
		SpecialBuild     string `json:"SpecialBuild"`
	} `json:"StringFileInfo"`
	VarFileInfo struct {
		Translation struct {
			LangID    string `json:"LangID"`
			CharsetID string `json:"CharsetID"`
		} `json:"Translation"`
	} `json:"VarFileInfo"`
	IconPath     string `json:"IconPath"`
	ManifestPath string `json:"ManifestPath"`
}

var (
	fileVersionMajor int
	fileVersionMinor int
	fileVersionPatch int
	fileVersionBuild int
	prodVersionMajor int
	prodVersionMinor int
	prodVersionPatch int
	prodVersionBuild int
	fileFlagsMask    string
	fileFlags        string
	fileOS           string
	fileType         string
	fileSubType      string
	comments         string
	companyName      string
	fileDescription  string
	fileVersion      string
	internalName     string
	legalCopyright   string
	legalTrademarks  string
	originalFilename string
	privateBuild     string
	productName      string
	productVersion   string
	specialBuild     string
	langID           string
	charsetID        string
	iconPath         string
	manifestPath     string
	sixtyFour        bool
	example          bool
	outputFile       string
	outputSpecific   bool
)

func init() {
	flag.StringVar(&charsetID, "", "", "Charset for the output file")
	flag.StringVar(&comments, "", "", "Comments")
	flag.StringVar(&companyName, "", "", "Company name")
	flag.StringVar(&legalCopyright, "", "", "Legal copyright")
	flag.StringVar(&fileDescription, "", "", "Description of file")
	flag.BoolVar(&example, "", false, "Just dump an example versioninfo.json, defaults to false")
	flag.StringVar(&fileVersion, "", "", "File version for file info")
	flag.StringVar(&iconPath, "", "", "Icon file name")
	flag.StringVar(&internalName, "", "", "Internal name for file info")
	flag.StringVar(&manifestPath, "", "", "Path to manifest file (optional)")
	flag.StringVar(&outputFile, "", "", "Output file name")
	flag.BoolVar(&outputSpecific, "", false, "Outputs specific file for amd64 and is836, ignore '-o', defaults to false")
	flag.StringVar(&originalFilename, "", "", "Original filename for file info")
	flag.StringVar(&specialBuild, "", "", "Special build info for file info")
	flag.StringVar(&legalTrademarks, "", "", "Legal trademarks for file info")
	flag.StringVar(&langID, "", "", "Language ID for file info")
	flag.StringVar(&charsetID, "", "", "Charset ID for file info")
	flag.BoolVar(&sixtyFour, "", false, "Output file for 64bit archetecture, default is FALSE")
	flag.IntVar(&fileVersionMajor, "", 1, "Major version of output file")
	flag.IntVar(&fileVersionMinor, "", 1, "Minor version of output file")
	flag.IntVar(&fileVersionPatch, "", 3, "Patch version of output file")
	flag.IntVar(&fileVersionBuild, "", 2, "Build # of output file")
	flag.IntVar(&prodVersionMajor, "", 1, "Major version of output product")
	flag.IntVar(&prodVersionMinor, "", 1, "Minor version of output product")
	flag.IntVar(&prodVersionPatch, "", 3, "Patch version of output product")
	flag.IntVar(&prodVersionBuild, "", 2, "Build # of product")
}

func main() {

}
