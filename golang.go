// Aplicação que exibe a hora atual, implicando que em cada atualização você terá uma saída diferente.
package mai
 import (
   "fmt"
   "net/http"
   "time" // precisa para exibir a hora
 )
 func main() {
     http.HandleFunc("/", handler)
     log.Println("listening...")
     err := http.ListenAndServe(":8080", nil)
     if err != nil {
         panic(err)
     }
 }
 func handler(w http.ResponseWriter, r *http.Request) {
     fmt.Fprintf(w, "Hello. The time is : " + time.Now().Format(time.RFC850))
 }
 
 ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// traz todo codigo fonte de uma pagina 
package main
import (
	"net/http"
	"io/ioutil"
	"os"
	"fmt"
)
func main() {
	for _, url := range os.Args[1:]{
		resp, err := http.Get(url)
		if err != nil{
			os.Exit(1)
		}
	
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil{
			os.Exit(1)
		}
		resp.Body.Close()
		
		fmt.Printf("%s", b)
	}
}
go build main.go
./main http://google.com.br

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// remove um item do array 
for indiceArray, item := range cadastros {
	if item.ID == params["id"] {
	    cadastros = append(cadastros[:indiceArray], cadastros[indiceArray+1:]...)
	    break
	}
}

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Lê todos os arquivos/pastas do diretorio atual e salva em um Slice 
package main    
import (
    "fmt"
    "path/filepath"
)

func main() {
    files, _ := filepath.Glob("*")
    fmt.Println(files) // traz todos os arquivos e pastas do diretorio atual
    fmt.Println(files[1]) // traz o arquivo/pasta com indice 2 do slice
}

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Lê todos os arquivos/pastas de um diretorio e o tamanho em bytes
package main
import (
    "fmt"
    "io/ioutil"
)
func main() {
    files, _ := ioutil.ReadDir("./")
    //files, _ := ioutil.ReadDir("C:/Users/")
    for _, f := range files {
            //fmt.Println(f.Name())
	    fmt.Println(f.Name(), f.Size(), "bytes")
    }
}

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Traz as informações somente dos arquivos em um diretorio
package main
import (
    "fmt"
    "os"
    "path/filepath"
)
func main() {
    dirname := "C:/" + string(filepath.Separator)
    d, err := os.Open(dirname)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer d.Close()
    fi, err := d.Readdir(-1)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    for _, fi := range fi {
        if fi.Mode().IsRegular() {
            fmt.Println(fi.Name(), fi.Size(), "bytes")
        }
    }
}

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Lê todos os arquivos/pastas de forma recursiva.
package main
import (
    "fmt"
    "os"
    "path/filepath"
)
func main() {
    searchDir := "D:/Projects"
    fileList := []string{}
    err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
        fileList = append(fileList, path)
        return nil
    })
    if err != nil{
        fmt.Println(err)
    }
    for _, file := range fileList {
        fmt.Println(file)
    }
}

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Move de um diretorio para outro
package main
   import (
       "fmt"
       "os"
   )
   func main() {
       err :=  os.Rename("/dir1/file1", "/dir2/file2")
       if err != nil {
           fmt.Println(err)
           return
       }
}
