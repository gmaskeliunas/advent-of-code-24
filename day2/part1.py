def main():
    count = 0
    reports = []
    with open("input.txt") as f:
        for line in f:
            reports.append([int(x) for x in line.split(sep=" ")])
    print(reports)

if __name__ == "__main__":
    main()