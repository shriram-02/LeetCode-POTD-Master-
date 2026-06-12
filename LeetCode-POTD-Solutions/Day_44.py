def main():
    N, T = map(int, input().split())  # number of rooms, total hours

    for _ in range(N):
        K, E = map(int, input().split())  # number of required intervals, energy limit

        required = [False] * T

        for _ in range(K):
            L, R = map(int, input().split())
            for t in range(L - 1, R):
                required[t] = True

        schedule = [0] * T
        energy_used = 0

        # Step 1: Assign Eco mode (1) to required hours
        for i in range(T):
            if required[i]:
                schedule[i] = 1
                energy_used += 1

        # Step 2: Try converting Eco (1) to Cool (2) (though it increases energy)
        if energy_used > E:
            diff = energy_used - E
            for i in reversed(range(T)):
                if schedule[i] == 1 and diff > 0:
                    schedule[i] = 2
                    energy_used += 1  # +2 instead of 1 = +1 extra
                    diff += 1         # makes it worse, but retained to match original logic

        if energy_used > E:
            print("NOT POSSIBLE")
        else:
            print(" ".join(map(str, schedule)))

if __name__ == "__main__":
    main()