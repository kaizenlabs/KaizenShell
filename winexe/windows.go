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
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Version struct {
	Major int `json:"Major"`
	Minor int `json:"Minor"`
	Patch int `json:"Patch"`
	Build int `json:"Build"`
}

type FixedFileInfo struct {
	FileVersion    Version `json:"FileVersion"`
	ProductVersion Version `json:"ProductVersion"`
	FileFlagsMask  string  `json:"FileFlagsMask"`
	FileFlags      string  `json:"FileFlags "`
	FileOS         string  `json:"FileOS"`
	FileType       string  `json:"FileType"`
	FileSubType    string  `json:"FileSubType"`
}

type StringFileInfo struct {
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
}

type Translation struct {
	LangID    string `json:"LangID"`
	CharsetID string `json:"CharsetID"`
}

type VersionInfo struct {
	FixedFileInfo  FixedFileInfo  `json:"FixedFileInfo"`
	StringFileInfo StringFileInfo `json:"StringFileInfo"`
	VarFileInfo    VarFileInfo    `json:"VarFileInfo"`
	IconPath       string         `json:"IconPath"`
	ManifestPath   string         `json:"ManifestPath"`
}

type VarFileInfo struct {
	Translation Translation `json:"Translation"`
}

// VersionInfo is a struct representing the final output file structure in JSON format

var (
	fileVersionMajor int
	fileVersionMinor int
	fileVersionPatch int
	fileVersionBuild int
	prodVersionMajor int
	prodVersionMinor int
	prodVersionPatch int
	prodVersionBuild int
	fileFlagsMask    string = "3f"
	fileFlags        string = "00"
	fileOS           string = "040004"
	fileType         string = "01"
	fileSubType      string = "00"
	comments         string
	companyName      string
	fileDescription  string
	fileVer          string
	internalName     string
	legalCopyright   string
	legalTrademarks  string
	originalFilename string
	privateBuild     string
	productName      string
	productVersion   string = "v1.0.0.0"
	specialBuild     string
	langID           string = "0409"
	charsetID        string = "04B0"
	iconPath         string
	manifestPath     string
	sixtyFour        bool
	example          bool
	outputFile       string = "./dist/windows/Homebrew.exe"
	outputSpecific   bool
)

func init() {
	// flag.StringVar(&charsetID, "", "", "Charset for the output file")
	flag.StringVar(&comments, "comments", "", "Comments")
	flag.StringVar(&companyName, "companyName", "", "Company name")
	flag.StringVar(&legalCopyright, "copyright", "", "Legal copyright")
	flag.StringVar(&fileDescription, "appName", "", "Name of the file")
	// flag.BoolVar(&example, "", false, "Just dump an example versioninfo.json, defaults to false")
	// flag.StringVar(&fileVer, "", "", "File version for file info")
	flag.StringVar(&iconPath, "icon", "", "Filepath for the app icon")
	// flag.StringVar(&internalName, "", "", "Internal name for file info")
	// flag.StringVar(&manifestPath, "", "", "Path to manifest file (optional)")
	// flag.StringVar(&outputFile, "", "", "Output file name")
	// flag.BoolVar(&outputSpecific, "", false, "Outputs specific file for amd64 and is836, ignore '-o', defaults to false")
	// flag.StringVar(&originalFilename, "", "", "Original filename for file info")
	// flag.StringVar(&specialBuild, "", "", "Special build info for file info")
	flag.StringVar(&legalTrademarks, "trademark", "", "Legal trademarks for file info")
	// flag.StringVar(&langID, "", "", "Language ID for file info")
	// flag.StringVar(&charsetID, "", "", "Charset ID for file info")
	flag.BoolVar(&sixtyFour, "sixtyFour", false, "Output file for 64bit archetecture, default is FALSE")
	flag.IntVar(&fileVersionMajor, "major", 1, "Major version of output file")
	flag.IntVar(&fileVersionMinor, "minor", 0, "Minor version of output file")
	flag.IntVar(&fileVersionPatch, "patch", 0, "Patch version of output file")
	flag.IntVar(&fileVersionBuild, "build", 0, "Build # of output file")
	flag.Parse()
}

func main() {

	f := FixedFileInfo{
		FileVersion: Version{
			Major: fileVersionMajor,
			Minor: fileVersionMinor,
			Patch: fileVersionPatch,
			Build: fileVersionBuild,
		},
		ProductVersion: Version{
			Major: fileVersionMajor,
			Minor: fileVersionMinor,
			Patch: fileVersionPatch,
			Build: fileVersionBuild,
		},
		FileFlagsMask: fileFlagsMask,
		FileFlags:     fileFlags,
		FileOS:        fileOS,
		FileType:      fileType,
		FileSubType:   fileSubType,
	}

	version := "v" + strconv.Itoa(fileVersionMajor) + "." + strconv.Itoa(fileVersionMinor) + "." + strconv.Itoa(fileVersionPatch) + "." + strconv.Itoa(fileVersionBuild)

	s := StringFileInfo{
		Comments:         comments,
		CompanyName:      companyName,
		FileDescription:  fileDescription,
		FileVersion:      version,
		InternalName:     fileDescription,
		LegalCopyright:   legalCopyright,
		LegalTrademarks:  legalTrademarks,
		OriginalFilename: fileDescription,
		PrivateBuild:     version,
		ProductName:      fileDescription,
		ProductVersion:   version,
		SpecialBuild:     version,
	}

	vfi := VarFileInfo{
		Translation{
			LangID:    langID,
			CharsetID: charsetID,
		},
	}

	v := VersionInfo{
		FixedFileInfo:  f,
		StringFileInfo: s,
		VarFileInfo:    vfi,
		IconPath:       iconPath,
	}

	toJSON, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("Erroring marshalling json versioninfo")
		os.Exit(1)
	}

	file, err := os.Create("./versioninfo.json")
	if err != nil {
		fmt.Println("Error creating file")
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write(toJSON)
	if err != nil {
		fmt.Println("Erroring marshalling json versioninfo")
		fmt.Println(err)
		os.Exit(1)
	}
}
