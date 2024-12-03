def main():
    left_location_ids = []
    right_location_ids = []

    with open("day1/input.txt") as f:
        for line in f:
            left_location_ids.append(int(line.split(sep="   ")[0]))
            right_location_ids.append(int(line.split(sep="   ")[1]))
    left_location_ids = sorted(left_location_ids)
    print(left_location_ids)
    right_location_ids = sorted(right_location_ids)
    print(right_location_ids)

    temp = []

    for x, y in zip(left_location_ids, right_location_ids):
        temp.append(abs(x-y))

    print(sum(temp))    


if __name__ == "__main__":
    main()