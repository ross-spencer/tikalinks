package main

import (
   "os"
   "fmt"
)

//callback for walk needs to match the following:
//type WalkFunc func(path string, info os.FileInfo, err error) error
func readFile (path string, fi os.FileInfo, err error) error {   
   fp := openFile(path)
   switch mode := fi.Mode(); {
   case mode.IsRegular():
      content, err := getFileContent(fp, fi)
      if err != nil {
         return err
      }     
      edat := getEntityData(content, fi.Name()) 
      collateEntities(edat)
   case mode.IsDir():
      fmt.Fprintln(os.Stderr, "INFO:", fi.Name(), "is a directory.")      
   default: 
      fmt.Fprintln(os.Stderr, "INFO: Something completely different.")
   }
   return nil
}

