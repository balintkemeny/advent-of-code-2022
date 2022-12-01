caloriesInventory = []
elfSum = 0

with open("input.txt", "r") as f:
    for line in f.readlines():
        if line == "\n":
            caloriesInventory.append(elfSum)
            elfSum = 0
        else:
            elfSum += int(line.rstrip())

if elfSum > 0:
    caloriesInventory.append(elfSum)

print(f"The elf carrying the most energy carries: {max(caloriesInventory)} calories.")

caloriesInventory.sort()
topThreeTotal = sum(caloriesInventory[-3:])
print(f"The total calories carried by the top 3 elves is: {topThreeTotal}.")