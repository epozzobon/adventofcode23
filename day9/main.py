s0, s1 = 0, 0
for line in open("day9/input.txt", "r"):
    nums = [int(p.strip()) for p in line.split(" ") if p.strip() != ""]
    derivatives = [nums]
    for _ in range(len(nums)):
        d = derivatives[-1]
        derivatives.append([d[i+1] - d[i] for i in range(len(d)-1)])
    for i in range(len(nums)-2, -1, -1):
        predictedNext = derivatives[i+1][-1] + derivatives[i][-1]
        predictedPrev = derivatives[i][0] - derivatives[i+1][0]
        derivatives[i] = [predictedPrev] + derivatives[i] + [predictedNext]
        print("   " * i, ("%5d " * len(derivatives[i])) % tuple(derivatives[i]))
    s0 += derivatives[0][0]
    s1 += derivatives[0][-1]
print(s0, s1)
