let input = Utils.read_whole_file "./lib/day_04/input.txt" |> String.split_on_char '\n' 

let rec has_dup = function
    | [] -> false
    | hd :: tl -> List.exists (( = ) hd) tl || has_dup tl

let rec unique_count = function
  | [] -> 0
  | hd :: tl -> 
      (if hd |> String.split_on_char ' ' |> has_dup then 0 else 1) + (unique_count tl)


let rec valid_count = function
  | [] -> 0
  | hd :: tl -> 
      let sorted_words = hd 
        |> String.split_on_char ' ' 
        |> List.map (fun w -> w 
          |> String.to_seq 
          |> List.of_seq 
          |> List.sort Char.compare 
          |> List.to_seq 
          |> String.of_seq) in
      
      (if sorted_words |> has_dup then 0 else 1) + (valid_count tl)

let part1 = string_of_int @@ unique_count input
let part2 = string_of_int @@ valid_count input
