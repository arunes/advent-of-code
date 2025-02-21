defmodule Day01 do
  defp get_numbers() do
    lines =
      File.stream!("./lib/day-01/input.txt")
      |> Stream.map(&String.trim_trailing/1)
      |> Enum.to_list()

    Enum.map(lines, fn ln -> Integer.parse(ln) |> elem(0) end)
  end

  def part1() do
    numbers = get_numbers()
    result = Enum.reduce(numbers, 0, fn n, acc -> n + acc end)
    IO.puts(result)
  end

  def part2() do
    numbers = get_numbers()
    result = part2_solve(numbers, 0, [])
    IO.puts(result)
  end

  def part2_solve([head | tail], acc, history) do
    total = acc + head

    if Enum.member?(history, total) do
      total
    else
      part2_solve(tail, total, [total | history])
    end
  end

  def part2_solve([], acc, history) do
    part2_solve(get_numbers(), acc, history)
  end
end
