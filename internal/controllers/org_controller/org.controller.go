package orgcontrollers

import (
	"fmt"

	"github.com/AdelGann/z0-backend-v1/internal/inputs/org_inputs"
	"github.com/AdelGann/z0-backend-v1/internal/services/org_service"
	"github.com/AdelGann/z0-backend-v1/pkg/utils/helpers/validations"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SaveOrg() {

}
func JoinOrg(c *fiber.Ctx) error {
	input := new(orginputs.JoinOrgInput)
	claims, err := validations.ExtractClaims(c)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := c.BodyParser(input); err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(fiber.Map{"error": "Failed while parsing the body"})
	}
	IdClaim, ok := claims["sub"].(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid email type in claims"})
	}
	res, err := orgservices.JoinOrg(input.Code, uuid.MustParse(IdClaim))
	if err != nil {
		fmt.Print(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"code": res, "state": fiber.StatusCreated})
}
func Invite(c *fiber.Ctx) error {
	invitation := new(orginputs.InviteOrgInput)
	claims, err := validations.ExtractClaims(c)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := c.BodyParser(invitation); err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(fiber.Map{"error": "Failed while parsing the body"})
	}
	IdClaim, ok := claims["sub"].(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid email type in claims"})
	}
	res, err := orgservices.SendInvitation(invitation.OrgID, invitation.Email, uuid.MustParse(IdClaim))

	if err != nil {
		fmt.Print(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"code": res, "state": fiber.StatusCreated})
}
