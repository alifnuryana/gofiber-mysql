package services

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gofiber-mysql/database"
	"gofiber-mysql/entity"
	"log"
)
import _ "gofiber-mysql/database"

func GetStudents(ctx *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT id, nama, npm, prodi FROM mahasiswa ORDER BY id")
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}
	defer rows.Close()
	var result entity.Students
	for rows.Next() {
		var student entity.Student
		if err := rows.Scan(&student.ID, &student.Nama, &student.NPM, &student.Prodi); err != nil {
			log.Fatal(err)
		}
		result.Students = append(result.Students, student)
	}
	return ctx.JSON(fiber.Map{
		"status":  "success",
		"student": result.Students,
	})
}

func PostStudent(ctx *fiber.Ctx) error {
	student := new(entity.Student)

	if err := ctx.BodyParser(student); err != nil {
		log.Fatal(err)
	}
	result, err := database.DB.Query("INSERT INTO mahasiswa (nama, npm, prodi) VALUES (?, ?, ?)", student.Nama, student.NPM, student.Prodi)
	if err != nil {
		return err
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	return ctx.JSON(fiber.Map{
		"status":  "success",
		"student": student,
	})
}

func PutStudent(ctx *fiber.Ctx) error {
	student := new(entity.Student)
	if err := ctx.BodyParser(student); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}
	res, err := database.DB.Query("UPDATE mahasiswa SET nama=?, npm=?, prodi=? WHERE id = ?", student.Nama, student.NPM, student.Prodi, student.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}
	fmt.Println(res)
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"student": student,
	})
}

func DeleteStudent(ctx *fiber.Ctx) error {
	student := new(entity.Student)

	if err := ctx.BodyParser(student); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	res, err := database.DB.Query("DELETE FROM mahasiswa WHERE id = ?", student.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}
	fmt.Println(res)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"student": student,
	})
}
