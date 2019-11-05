package helper

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func FileToBase64(path string) (string, error) {
	_, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buff), nil
}

func FileInStringChanger(str string) (string, error) {
	rst := str
	pattern := `@{(\.|/|[a-zA-Z0-9]+)?:?/[a-zA-Z0-9/./_/-]+(/$)?}`
	reg := regexp.MustCompile(pattern)
	target := reg.FindAllString(str, -1)

	for _, v := range target {
		val, err := FileToBase64(v[2 : len(v)-1])
		if err != nil {
			return "", err
		}
		rst = strings.Replace(rst, v, "\""+val+"\"", -1)
	}
	return rst, nil
}

//获取一个目录下所有文件的名字（未限制后缀）
func GetPathFileName(path string) []string {
	var ret []string
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if !file.IsDir() {
			ret = append(ret, path+"/"+file.Name())
		}
	}
	return ret
}

//遍历文件夹及子文件夹下的所有文件--只输出文件名=file_name的文件,file_name=""则表示输出所有
//func GetAllFile(pathname string,s map[string][]string,list_path []string,file_name string) ( map[string][]string, error) {
//	//var list_path []string
//	rd, err := ioutil.ReadDir(pathname)
//	if err != nil {
//		fmt.Println("read dir fail:", err)
//		return s, err
//	}
//	for _, fi := range rd {
//		if fi.IsDir() {
//			fullDir := pathname + "/" + fi.Name()
//			s, err = GetAllFile(fullDir,s,file_name)
//			if err != nil {
//				fmt.Println("read dir fail:", err)
//				return s, err
//			}
//		} else {
//			if fi.Name()=="images.list"{
//				list_path= append(list_path,pathname)
//				//fmt.Println(pathname)
//			}
//			if file_name==fi.Name(){
//				//fullName := pathname + "/" + fi.Name()
//			}
//if file_name==""{
//	if fi.Name()==file_name{
//		fullName := pathname + "/" + fi.Name()
//		s[pathname] =
//		s = append(s, fullName)
//	}
//}else {
//	fullName := pathname + "/" + fi.Name()
//	s = append(s, fullName)
//}
//		}
//	}
//	fmt.Println(list_path)
//	return s, nil
//}
