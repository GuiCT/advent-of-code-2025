# Wait a minute. Intruder alert!
import os
import argparse
import re
import numpy as np
from numpy.typing import NDArray
from ortools.linear_solver import pywraplp

DESIRED_STATE_REGEX = re.compile(r"[(\.\#+)]")
BUTTON_REGEX = re.compile(r"\((.+?)\)")
DESIRED_JOLTAGE_REGEX = re.compile(r"{(.+)}")


class JoltageMachine:
    desired_state: str
    buttons: NDArray[np.int64]
    joltage: NDArray[np.int64]

    def __init__(self, line: str):
        match_state = DESIRED_STATE_REGEX.search(line)
        if match_state is None:
            raise ValueError("Could not find desired state in line")
        self.desired_state = match_state.group(0)
        match_joltage = DESIRED_JOLTAGE_REGEX.search(line)
        if match_joltage is None:
            raise ValueError("Could not find desired joltage in line")
        joltage_str = match_joltage.group(1)
        joltage_nums = joltage_str.split(",")
        joltage = np.array([int(num) for num in joltage_nums], dtype=np.int64)
        self.joltage = joltage
        iter_buttons = list(BUTTON_REGEX.finditer(line))
        amount_buttons = len(iter_buttons)
        buttons = np.zeros((amount_buttons, len(self.joltage)), dtype=np.int64)
        for button_idx, match in enumerate(iter_buttons):
            button_str = match.group(1)
            nums = button_str.split(",")
            for num in nums:
                num_int = int(num)
                buttons[button_idx, num_int] = 1
        self.buttons = buttons

    def find_part2_solution(self) -> int:
        solver = pywraplp.Solver.CreateSolver("SAT")
        if solver is None:
            raise RuntimeError("Could not create solver")
        # Variables are the amount of times each button is pressed
        button_vars = []
        for i in range(self.buttons.shape[0]):
            var = solver.NumVar(0, solver.infinity(), f"button_{i}_presses")
            button_vars.append(var)
        # Constraints is the joltage that must be met
        for i in range(self.joltage.size):
            constraint = solver.RowConstraint(
                float(self.joltage[i]),
                float(self.joltage[i]),
                f"joltage_constraint_{i}",
            )
            for j in range(self.buttons.shape[0]):
                constraint.SetCoefficient(button_vars[j], float(self.buttons[j, i]))
        # Objective is to minimize the total button presses
        objective = solver.Objective()
        for var in button_vars:
            objective.SetCoefficient(var, 1)
        objective.SetMinimization()
        status = solver.Solve()
        if status != pywraplp.Solver.OPTIMAL:
            raise RuntimeError("Solver did not find optimal solution")
        total_presses = int(objective.Value())
        return total_presses


def get_lines(use_example: bool) -> list[str]:
    folder_name = "examples" if use_example else "inputs"
    file_path = os.path.join(folder_name, "day10.txt")
    with open(file_path) as f:
        return [line.strip() for line in f.readlines()]


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    _ = parser.add_argument(
        "-e",
        "--example",
        action="store_true",
        help="Use example input instead of puzzle input",
    )
    args = parser.parse_args()
    lines = get_lines(args.example)
    part2_solution_sum = 0
    for line in lines:
        machine = JoltageMachine(line)
        part2_solution = machine.find_part2_solution()
        part2_solution_sum += part2_solution
    print(f"Part 2 Solution Sum: {part2_solution_sum}")
