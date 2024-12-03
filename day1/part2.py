def main():
    left_location_ids = []
    right_location_ids = []

    with open("input.txt") as f:
        for line in f:
            left_location_ids.append(int(line.split(sep="   ")[0]))
            right_location_ids.append(int(line.split(sep="   ")[1]))

    temp = []

    for number in left_location_ids:
        if number in right_location_ids:
            count = [x for x in right_location_ids if x == number]
            temp.append(number*len(count))

    print(sum(temp))    


if __name__ == "__main__":
    main()