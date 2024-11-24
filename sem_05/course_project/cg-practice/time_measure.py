import matplotlib
import matplotlib.pyplot as plt

try:
    with open("./time.txt", "r", encoding = "utf-8") as f:
        depth, time = [], []
        all_file = f.read()
        lines = all_file.split("\n")
        for line in lines:
            data = line.split(" ")
            if len(data) == 2:
                depth.append(int(data[0]))
                time.append(float(data[1]))

        plt.title("Скорость работы алгоритма трассировки лучей")
        plt.xlabel("Глубина рекурсии")
        plt.ylabel("Время работы, мс")
        plt.grid(True)
        
        plt.plot(depth, time)

        plt.savefig("graph.svg")

except Exception as e:
    print(e)
    input()

