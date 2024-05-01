package main

import (
	"context"
	"fmt"
	"time"

	"github.com/evertras/bubble-table/table"
	"github.com/jackc/pgx/v5"
)

// helper functions
func getRecipe() []Recipe {
	query, err := conn.Query(context.Background(), "select * from recipe")
	if err != nil {
		fmt.Println(err)
	}
	rows, err := pgx.CollectRows(query, pgx.RowToStructByName[Recipe])
	if err != nil {
		fmt.Println(err)
	}
	defer query.Close()
	return rows
}

func getAllRecipesByName() table.Model {
	entries := getRecipe()
	length := len("Name")
	rows := []table.Row{}
	for _, o := range entries {
		length = max(length, len(o.Name))
		rows = append(rows, table.NewRow(table.RowData{
			"Name": o.Name,
		}))
	}
	columns := []table.Column{
		table.NewColumn("Name", "Name", length),
	}
	tbl := table.New(columns).WithRows(rows)
	return tbl
}

func getRecipeByName(prompt string) table.Model {

	var res Recipe
	query := fmt.Sprintf("SELECT * FROM recipe WHERE LOWER(name) LIKE LOWER('%%%s%%')", prompt)
	err := conn.QueryRow(context.Background(), query).Scan(&res.Recipeid, &res.Name, &res.Description, &res.Instructions, &res.Preptime, &res.Cooktime, &res.Totaltime)
	length := lenRecipe{
		lenid:    8,
		lenname:  len(res.Name),
		lendesc:  len(res.Description),
		lenins:   len(res.Instructions),
		lenprep:  8,
		lencook:  8,
		lentotal: 9,
	}
	if err != nil {
		fmt.Println("err: ", err, prompt)
	}

	rows := []table.Row{table.NewRow(table.RowData{
		"Recipeid":     res.Recipeid,
		"Name":         res.Name,
		"Description":  res.Description,
		"Instructions": res.Instructions,
		"Preptime":     res.Preptime.Format(time.TimeOnly),
		"Cooktime":     res.Cooktime.Format(time.TimeOnly),
		"Totaltime":    res.Totaltime.Format(time.TimeOnly),
	})}
	columns := []table.Column{
		table.NewColumn("Recipeid", "Recipeid", length.lenid),
		table.NewColumn("Name", "Name", length.lenname),
		table.NewColumn("Description", "Description", length.lendesc),
		table.NewColumn("Instructions", "Instructions", length.lenins),
		table.NewColumn("Preptime", "Preptime", length.lenprep),
		table.NewColumn("Cooktime", "Cooktime", length.lencook),
		table.NewColumn("Totaltime", "Totaltime", length.lentotal),
	}
	tbl := table.New(columns).WithRows(rows)
	return tbl
}

// func gethRecipeByIngredient() table.Model {

// }
// func getFavList() table.Model {

// }
// func getPopularRecipes() table.Model {

// }
