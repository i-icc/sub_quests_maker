def txt_to_sql(file_name):
    with open(file_name, mode="r") as f:
        lines = f.readlines()
    
    table_name = lines[0].rstrip()
    result = [f"INSERT INTO {table_name}(instruction) VALUES"]
    for c in lines[1:]:
        c = c.rsplit()[0]
        if c == "end" or c == "":
            break
        result.append(f'("{c}"),')
    result[-1] = result[-1][:-1]+";"
    result.append("\n")
    with open("../insert.sql", mode="a") as f:
        f.writelines(result)


def main():
    for i in range(4):
        txt_to_sql(f"{i}.txt")


if __name__ == "__main__":
    main()
