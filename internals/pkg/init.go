package pkg

var pkgObjects *PkgWrapper

func Init(pkgWrapper *PkgModules) (*PkgWrapper, error) {
	if pkgObjects != nil {
		return pkgObjects, nil
	}

	allPkg := new(PkgWrapper)
	allPkg.Dialogflow = pkgWrapper
	allPkg.Telegram = pkgWrapper

	pkgObjects = allPkg

	return pkgObjects, nil
}
