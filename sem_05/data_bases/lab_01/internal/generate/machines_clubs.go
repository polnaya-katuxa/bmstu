package generate

import (
	"fmt"
	"git.parasha.space/go/libs/faker"
)

type MachineInClub struct {
	ID        int `csv:"id"`
	ClubID    int `csv:"id_club"`
	MachineID int `csv:"id_machine"`
}

func machineInClub(machines []Machine, clubs []ComputerClub) MachineInClub {
	i := faker.IntWithLimits(0, len(machines))
	j := faker.IntWithLimits(0, len(clubs))

	m := MachineInClub{
		ClubID:    j + 1,
		MachineID: i + 1,
	}

	return m
}

func MachinesInClubs(n int, machines []Machine, clubs []ComputerClub) []MachineInClub {
	machinesInClubs := make([]MachineInClub, n)

	for i := range machinesInClubs {
		machinesInClubs[i] = machineInClub(machines, clubs)
		machinesInClubs[i].ID = i + 1
		fmt.Printf("machines in clubs: %d/%d\r", i, n)
	}

	return machinesInClubs
}
