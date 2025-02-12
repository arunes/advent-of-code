let input = Utils.read_whole_file "./lib/day_02/input.txt" |> String.split_on_char '\n'

let checksum_row row = 
  let values = row |> String.split_on_char '\t' |> List.map (fun i -> int_of_string i) in
  let max = List.fold_left max 0 values in
  let min = List.fold_left min (List.nth values 0) values in
  max - min

let rec checksum list = 
  match list with
  | [] -> 0
  | head :: [] -> checksum_row head
  | head :: tail -> 
    checksum_row head + checksum tail

let evenly_divisible_row row = 
  let rec calc values all = match values with
  | [] -> 0
  | head :: [] ->
      let canditates = List.filter (fun i ->
        i <> head && i < head && float_of_int head /. float_of_int i = float_of_int (head / i)) all in
      if List.length canditates = 0 then 0 else head / (List.hd canditates)
  | head :: tail -> 
      let canditates = List.filter (fun i ->
        i <> head && i < head && float_of_int head /. float_of_int i = float_of_int (head / i)) all in
      if List.length canditates = 0 then calc tail all else head / (List.hd canditates)
  in

  let values = row |> String.split_on_char '\t' |> List.map (fun i -> int_of_string i) in
  calc values values

let rec evenly_divisible list = 
  match list with
  | [] -> 0
  | head :: [] -> evenly_divisible_row head
  | head :: tail -> 
      evenly_divisible_row head + evenly_divisible tail

let part1 = string_of_int (checksum input)
let part2 = string_of_int (evenly_divisible input)
