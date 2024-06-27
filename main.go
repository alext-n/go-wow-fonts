package main

import (
	"fmt"
	"os"
)

func main() {
	filenames := [22]string{
		"2002.ttf",
		"2002B.ttf",
		"615970.slug",
		"615974.slug",
		"ARHei.ttf",
		"arheiuhk_bd.ttf",
		"ARIALN.ttf",
		"ARKai_C.ttf",
		"ARKai_T.ttf",
		"bHE100M.ttf",
		"bHE101B.ttf",
		"bKAI00M.ttf",
		"blei00d.ttf",
		"FRIZQT_.ttf",
		"FRIZQT_CYR.ttf",
		"K_Damage.ttf",
		"K_Pagetext.ttf",
		"LICENSE.txt",
		"MORPHEUS.ttf",
		"MORPHEUS_CYR.ttf",
		"skurri.ttf",
		"SKURRI_CYR.ttf",
	}

	data, err := os.ReadFile("input.ttf")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, filename := range filenames {
		// Create a new file.
		f, err := os.Create(fmt.Sprint("outputs/", filename))
		if err != nil {
			fmt.Println(err)
			return
		}

		// Write the data to the new file.
		_, err = f.Write(data)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Close the file.
		if err := f.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}
}