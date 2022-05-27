package main

import (
  "os"
  "time"
  "math/rand"
  "strings"
  "errors"
  "fmt"
  "log"
  "image"
  "image/png"
  "image/color"
  "image/draw"
  "encoding/json"
  "io/ioutil"
)

func trim_string(trait string) string {
  spstr := strings.Split(trait,"_")
  restr := ""
  if len(spstr) == 1 {
    restr = spstr[0]
  }
  if len(spstr) > 1 {
    for i, k := range spstr {
      restr+=k
      if i != len(spstr)-1 {
        restr+= " "
      }
    }
  }
  return restr
}

func get_number(w string,b int) string {
  a := ""
  if b < 10 {
    a = fmt.Sprintf("%s #000%d",w,b)
  }
  if b > 9 && b < 100 {
    a = fmt.Sprintf("%s #00%d",w,b)
  }
  if b > 99 && b < 1000 {
    a = fmt.Sprintf("%s #0%d",w,b)
  }
  if b > 999 && b < 10000 {
    a = fmt.Sprintf("%s #%d",w,b)
  }
  return a
}

func faiLog(e error) {
  if e != nil {
    fmt.Println(e)
    log.Fatal("error!")
    panic(e)
  }
}

func genFromArray(arr []fileNames, len int) string {
  rand.Seed(time.Now().UnixNano())
  index := rand.Intn(len)
  return arr[index].name
}

type fileNames struct {name string}

func getFileNames(dir string) []fileNames {
  fileInfos, err := ioutil.ReadDir(dir)
  if err != nil {
    fmt.Println("Error in accessing directory:", err)
  }

  pngCounter := 0

  for _, k := range fileInfos {
    if strings.HasSuffix(k.Name(),".png") {
      pngCounter += 1
    }
  }

  fileInDir := make([]fileNames,0,pngCounter)

  for _, filename := range fileInfos {
    if strings.HasSuffix(filename.Name(),".png") {
      address:=strings.Join(strings.Split(filename.Name(),".png"),"")
      fileInDir = append(fileInDir,fileNames{address})
    }
  }

  return fileInDir
}

func isFolder(address string) {
  if _, err := os.Stat(address); errors.Is(err,os.ErrNotExist) {
    err := os.Mkdir(address,os.ModePerm)
    if err != nil {
      log.Println(err)
    }
  }
}

func fullAddressParentDir() string {
  path, err := os.Getwd()
  if err != nil {
    log.Println(err)
  }
  return path
}

func timeunix() int64 {
  now := time.Now().Unix()
  return now
}

type NFTTrait struct {
  Trait_type string `json:"trait_type"`
  Value string `json:"value"`
}

type NFTFile struct {
  Uri string `json:"uri"`
  Medium_type string `json:"type"`
}

type NFTCreator struct {
  Address string `json:"address"`
  Share int64 `json:"share"`
}

type NFTAuthor struct {
  Author_name string `json:"author_name"`
}

type NFTProperties struct {
  Files []NFTFile `json:"files"`
  Category string `json:"category"`
  Creators []NFTCreator `json:"creators"`
}

type NFTCollection struct {
  Name string `json:"name"`
  Family string `json:"family"`
}

type SolanaMetaplexNFTMetadata struct {
  Name string `json:"name"`
  Symbol string `json:"symbol"`
  Description string `json:"description"`
  Seller_fee_basis_points int64 `json:"seller_fee_basis_points"`
  ExternalURL string `json:"external_url"`
  Image string `json:"image"`
  Collection NFTCollection `json:"collection"`
  Properties NFTProperties `json:"properties"`
  Attributes []NFTTrait `json:"attributes"`
  Compiler string `json:"compiler"`
}

type ERC721NFTMetadata struct {
  Name string `json:"name"`
  Description string `json:"description"`
  Image string `json:"image"`
  Attributes []NFTTrait `json:"attributes"`
}

func main() {
  background := getFileNames("./background/")
  face := getFileNames("./face/")
  bottom := getFileNames("./bottom-clothing/")
  accessory := getFileNames("./accessories/")
  upper := getFileNames("./upper-clothing/")
  hair := getFileNames("./hair/")
  hands := getFileNames("./hands/")
  location := getFileNames("./location/")
  feet := getFileNames("./feet/")

  comb_710 := make([]string,0,3500)

  for len(comb_710) != 3500 {
    a := genFromArray(background,len(background))
    b := genFromArray(face,len(face))
    c := genFromArray(bottom,len(bottom))
    d := genFromArray(accessory,len(accessory))
    e := genFromArray(upper, len(upper))
    f := genFromArray(hair,len(hair))
    g := genFromArray(hands,len(hands))
    h := genFromArray(location,len(location))
    i := genFromArray(feet,len(feet))

    comb := a+"-"+b+"-"+c+"-"+d+"-"+e+"-"+f+"-"+g+"-"+h+"-"+i

    binary := false

    for _, x := range comb_710 {
      if x == comb {
        binary = true
        break
      }
    }

    if binary == false {
      comb_710=append(comb_710,comb)
      fmt.Println(comb)
    }
  }

  isFolder("assets")

  for i, loop := range comb_710 {
    background_address := strings.Split(loop,"-")[0]+".png"
    background_open,err := os.Open(background_address)
    faiLog(err)
    defer background_open.Close()
    background_image,_,err := image.Decode(background_open)
    faiLog(err)

    face_address := strings.Split(loop,"-")[1]+".png"
    face_open,err := os.Open(face_address)
    faiLog(err)
    defer face_open.Close()
    face_image,_,err := image.Decode(face_open)
    faiLog(err)

    bottom_address := strings.Split(loop,"-")[2]+".png"
    bottom_open,err := os.Open(bottom_address)
    faiLog(err)
    defer bottom_open.Close()
    bottom_image,_,err := image.Decode(bottom_open)
    faiLog(err)

    accessory_address := strings.Split(loop,"-")[3]+".png"
    accessory_open,err := os.Open(accessory_address)
    faiLog(err)
    defer accessory_open.Close()
    accessory_image,_,err := image.Decode(accessory_open)
    faiLog(err)

    upper_address := strings.Split(loop,"-")[4]+".png"
    upper_open,err := os.Open(upper_address)
    faiLog(err)
    defer upper_open.Close()
    upper_image,_,err := image.Decode(upper_open)
    faiLog(err)

    hair_address := strings.Split(loop,"-")[5]+".png"
    hair_open,err := os.Open(hair_address)
    faiLog(err)
    defer hair_open.Close()
    hair_image,_,err := image.Decode(hair_open)
    faiLog(err)

    hands_address := strings.Split(loop,"-")[6]+".png"
    hands_open,err := os.Open(hands_address)
    faiLog(err)
    defer hands_open.Close()
    hands_image,_,err := image.Decode(hands_open)
    faiLog(err)

    location_address := strings.Split(loop,"-")[7]+".png"
    location_open,err := os.Open(location_address)
    faiLog(err)
    defer location_open.Close()
    location_image,_,err := image.Decode(location_open)
    faiLog(err)

    feet_address := strings.Split(loop,"-")[8]+".png"
    feet_open,err := os.Open(feet_address)
    faiLog(err)
    defer feet_open.Close()
    feet_image,_,err := image.Decode(feet_open)
    faiLog(err)

    bgImg := image.NewRGBA(image.Rect(0, 0, 372, 612))
    draw.Draw(bgImg, bgImg.Bounds(), &image.Uniform{color.RGBA{0,0,0,0}}, image.ZP, draw.Src)

    offset := image.Pt(0,0)

    draw.Draw(bgImg,background_image.Bounds().Add(offset),background_image,image.ZP,draw.Over)

    background_open.Close()

    draw.Draw(bgImg,face_image.Bounds().Add(offset),face_image,image.ZP,draw.Over)

    face_open.Close()

    draw.Draw(bgImg,location_image.Bounds().Add(offset),location_image,image.ZP,draw.Over)

    location_open.Close()

    draw.Draw(bgImg,hair_image.Bounds().Add(offset),hair_image,image.ZP,draw.Over)

    hair_open.Close()

    draw.Draw(bgImg,feet_image.Bounds().Add(offset),feet_image,image.ZP,draw.Over)

    feet_open.Close()

    draw.Draw(bgImg,bottom_image.Bounds().Add(offset),bottom_image,image.ZP,draw.Over)

    bottom_open.Close()

    draw.Draw(bgImg,upper_image.Bounds().Add(offset),upper_image,image.ZP,draw.Over)

    upper_open.Close()

    draw.Draw(bgImg,hands_image.Bounds().Add(offset),hands_image,image.ZP,draw.Over)

    hands_open.Close()

    draw.Draw(bgImg,accessory_image.Bounds().Add(offset),accessory_image,image.ZP,draw.Over)

    accessory_open.Close()

    path := fmt.Sprintf("assets/%d.png",i)
    f,_ := os.Create(path)
    png.Encode(f,bgImg)

    f.Close()

    solana_metaplex_meta := SolanaMetaplexNFTMetadata{
      Name: get_number("",i+1),
      Symbol: "",
      Description: "",
      Seller_fee_basis_points: 718,
      ExternalURL: "",
      Image: fmt.Sprintf("%d.png",i),
      Collection: NFTCollection{
        Name: "",
        Family: "",
      },
      Properties: NFTProperties{
        Files: []NFTFile{
          NFTFile{
            Uri: fmt.Sprintf("%d.png",i),
            Medium_type: "image/png",
          },
        },
        Category: "image",
        Creators: []NFTCreator{
          NFTCreator{
            Address: "",
            Share:100,
          },
        },
      },
      Attributes: []NFTTrait{
        NFTTrait{
          Trait_type: "background",
          Value: trim_string(strings.Split(loop,"-")[0]),
        },
        NFTTrait{
          Trait_type: "face",
          Value: trim_string(strings.Split(loop,"-")[1]),
        },
        NFTTrait{
          Trait_type: "bottom-clothing",
          Value: trim_string(strings.Split(loop,"-")[2]),
        },
        NFTTrait{
          Trait_type: "accessory",
          Value: trim_string(strings.Split(loop,"-")[3]),
        },
        NFTTrait{
          Trait_type: "upper-clothing",
          Value: trim_string(strings.Split(loop,"-")[4]),
        },
        NFTTrait{
          Trait_type: "hair",
          Value: trim_string(strings.Split(loop,"-")[5]),
        },
        NFTTrait{
          Trait_type: "hands",
          Value: trim_string(strings.Split(loop,"-")[6]),
        },
        NFTTrait{
          Trait_type: "location",
          Value: trim_string(strings.Split(loop,"-")[7]),
        },
        NFTTrait{
          Trait_type: "feet",
          Value: trim_string(strings.Split(loop,"-")[8]),
        },
      },
      Compiler: "Maddy Generator",
    }

    erc721_meta := ERC721NFTMetadata{
      Name: get_number("",i+1),
      Description: "",
      Image: fmt.Sprintf("%d.png",i),
      Attributes: []NFTTrait{
        NFTTrait{
          Trait_type: "background",
          Value: trim_string(strings.Split(loop,"-")[0]),
        },
        NFTTrait{
          Trait_type: "face",
          Value: trim_string(strings.Split(loop,"-")[1]),
        },
        NFTTrait{
          Trait_type: "bottom-clothing",
          Value: trim_string(strings.Split(loop,"-")[2]),
        },
        NFTTrait{
          Trait_type: "accessory",
          Value: trim_string(strings.Split(loop,"-")[3]),
        },
        NFTTrait{
          Trait_type: "upper-clothing",
          Value: trim_string(strings.Split(loop,"-")[4]),
        },
        NFTTrait{
          Trait_type: "hair",
          Value: trim_string(strings.Split(loop,"-")[5]),
        },
        NFTTrait{
          Trait_type: "hands",
          Value: trim_string(strings.Split(loop,"-")[6]),
        },
        NFTTrait{
          Trait_type: "location",
          Value: trim_string(strings.Split(loop,"-")[7]),
        },
        NFTTrait{
          Trait_type: "feet",
          Value: trim_string(strings.Split(loop,"-")[8]),
        },
      },
    }

    if i == -1 {fmt.Println(solana_metaplex_meta)} 

    content, err := json.Marshal(erc721_meta)

    faiLog(err)

    err = ioutil.WriteFile(fmt.Sprintf("assets/%d.json",i),content,0644)

    faiLog(err)

    fmt.Println(strings.Split(loop,"-"))
  }
}
