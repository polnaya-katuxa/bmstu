package generate

import (
	"fmt"
	"git.parasha.space/go/libs/faker"
)

type GameOnMachine struct {
	GameID          int `csv:"id_game"`
	MachineInClubID int `csv:"id_machines_clubs"`
}

func gameOnMachine(machinesInClubs []MachineInClub, games []Game) GameOnMachine {
	i := faker.IntWithLimits(0, len(machinesInClubs))
	j := faker.IntWithLimits(0, len(games))

	g := GameOnMachine{
		GameID:          j + 1,
		MachineInClubID: i + 1,
	}

	return g
}

func GamesOnMachines(n int, machinesInClubs []MachineInClub, games []Game) []GameOnMachine {
	gamesOnMachines := make([]GameOnMachine, n)

	gamesOnMachinesMap := make(map[GameOnMachine]struct{})

	for i := 0; i < n; i++ {
		gamesOnMachines[i] = gameOnMachine(machinesInClubs, games)
		if _, ok := gamesOnMachinesMap[gamesOnMachines[i]]; ok {
			i--
			continue
		} else {
			gamesOnMachinesMap[gamesOnMachines[i]] = struct{}{}
		}

		fmt.Printf("games on machines: %d/%d\r", i, n)
	}

	return gamesOnMachines
}
