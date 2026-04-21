defmodule PerfectSquares do
  def count_perfect_squares do
    n_str = IO.gets("") |> String.trim()
    {n, _} = Integer.parse(n_str)

    numbers_str = IO.gets("") |> String.trim()

    numbers =
      numbers_str
      |> String.split(" ")
      |> Enum.map(&String.to_integer/1)

      if length(numbers) != n do
      IO.puts("Number count mismatch")
      :ok

    else
      count =
        numbers
        |> Enum.count(fn number ->
          root = :math.sqrt(number) |> round()
          root * root == number
        end)

      IO.puts(count)
    end
  end
end

PerfectSquares.count_perfect_squares()
