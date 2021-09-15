package main

func main() {
	loadConfig()
	for _, v := range projectMap {
		version := getVersion(spliceDirStr(v.Path, v.Versionfile))
		// <yourObjectName>上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
		// <yourLocalFileName>由本地文件路径加文件名包括后缀组成，例如/users/local/myfile.txt。
		Admission(version)
		for _, file := range v.Files {
			uplaod(v.OssPathEndpoint, v.BucketName, spliceDirStr(spliceDirStr(v.OssPath, version), file), spliceDirStr(v.Path, file))
			writeVersion(version, file, spliceDirStr(v.Path, file))
		}
		Do(spliceDirStr(v.Path, "/version.txt"))
		uplaod(v.OssPathEndpoint, v.BucketName, spliceDirStr(spliceDirStr(v.OssPath, version), "/version.txt"), spliceDirStr(v.Path, "/version.txt"))
	}
}
