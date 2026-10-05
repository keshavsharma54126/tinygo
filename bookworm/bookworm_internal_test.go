package main

import (
	"fmt"
	"testing"
)

func TestLoadBookworms_Success(t *testing.T) {
	test_cases := map[string]struct{
		bookWormFileName string 
		want []Bookworms
		wantErr bool
	}{
	"file_exitsts":{
		bookworm_file:"./testdata/bookworm.json",
		want:[]Bookworm{{Name:"Fadi",
		Books:[]Books{
			Author:"Margaret Atwood",
			Title:"The Handmaid's Tale",
		}},
		{
			Name:"Peggy",
			Books:[]Books{
				Author: "Margaret Atwood",
				Title:"Oryx and Crake",
			},
		},
	
	},
		wantErr:false,
	},
	"file_not_exist":{bookworm_file:"./bookworm.json",
		want:[]Bookworm{{Name:"Fadi",
		Books:[]Books{
			Author:"Margaret Atwood",
			Title:"The Handmaid's Tale",
		}},
		{
			Name:"Peggy",
			Books:[]Books{
				Author: "Margaret Atwood",
				Title:"Oryx and Crake",
			},
		},
	
	},
		wantErr:false,
},
	"invalid_json":{bookworm_file:"./testdata/bookworm.json",
		want:[]Bookworm{{name:"Fadi",
		books:[]Books{
			author:"Margaret Atwood",
			title:"The Handmaid's Tale",
		}},
		{
			name:"Peggy",
			books:[]Books{
				author: "Margaret Atwood",
				title:"Oryx and Crake",
			},
		},
	
	},
		wantErr:false,
}
	}
	 
}



















