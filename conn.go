package main

import (
	"context"
	"fmt"

	"github.com/evertras/bubble-table/table"
	"github.com/jackc/pgx/v5"
)

// helper functions
func getRecipe() []recipe {
	query, err := conn.Query(context.Background(), "select * from recipe")
	if err != nil {
		fmt.Println(err)
	}
	rows, err := pgx.CollectRows(query, pgx.RowToStructByName[recipe])
	if err != nil {
		fmt.Println(err)
	}
	return rows
}

func getAllRecipesByName() table.Model {
	entries := getRecipe()
	columns := []table.Column{
		table.NewColumn("Name", "Name", len("Name")),
		table.NewColumn("Recipeid", "Recipeid", len("Recipeid")),
		table.NewColumn("Description", "Description", len("Description")),
		table.NewColumn("Instructions", "Instructions", len("Instructions")),
		table.NewColumn("Preptime", "Preptime", len("Preptime")),
		table.NewColumn("Cooktime", "Cooktime", len("Cooktime")),
		table.NewColumn("Totaltime", "Totaltime", len("Totaltime")),
	}
	rows := []table.Row{}
	for _, o := range entries {
		rows = append(rows, table.NewRow(table.RowData{
			"Name":         o.Name,
			"Recipeid":     o.Recipeid,
			"Description":  o.Description,
			"Instructions": o.Instructions,
			"Preptime":     o.Preptime,
			"Cooktime":     o.Cooktime,
			"Totaltime":    o.Totaltime,
		}))
	}
	tbl := table.New(columns).WithRows(rows)
	return tbl
}

// func getRecipeByName() table.Model {

// }
// func gethRecipeByIngredient() table.Model {

// }
// func getFavList() table.Model {

// }
// func getPopularRecipes() table.Model {

// }
