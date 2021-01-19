package meta

type FileMeta struct {
	FileSha1	string
	FileName	string
	FileSize	int64
	Location 	string
	UploadAt	string
}


var fileMetas map[string]FileMeta

func init()  {
	fileMetas = make(map[string]FileMeta)
}

func UpdateFileMeta(fmate FileMeta)  {
	fileMetas[fmate.FileSha1] = fmate
}

func GetFileMeta(filesha1 string) FileMeta {
	return fileMetas[filesha1]
}

//获取批量列表
func GetLastFileMetas(count int) []FileMeta {
	fmArray := make([]FileMeta, len(fileMetas))
	for _, v := range fileMetas {
		fmArray = append(fmArray, v)
	}
	//需要排序
	return fmArray[0:count]
}

func DeleteFileMeta(filesha1 string) {
	delete(fileMetas, filesha1)
}
