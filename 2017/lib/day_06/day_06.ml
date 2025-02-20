let input = [ 10; 3; 15; 10; 5; 15; 5; 15; 9; 2; 5; 8; 5; 2; 3; 6 ]

(* finds the max value in the list and returns its index *)
let get_max_pos lst =
  let rec calc current_pos max_pos max_value = function
    | [] -> max_pos
    | hd :: tl ->
        let new_max_pos = if hd > max_value then current_pos else max_pos in
        let new_max_value = if hd > max_value then hd else max_value in
        calc (current_pos + 1) new_max_pos new_max_value tl
  in
  calc 0 0 0 lst

let find_duplicate is_part_2 =
  let length = input |> List.length in
  let get_next_index idx = if idx + 1 = length then 0 else idx + 1 in
  let rec redistribute cycle remaining current_index current_list
      completed_lists =
    (* if remaining is zero and current list is in completed lists, return cycle, otherwise keep redistributing *)
    let found_dup_index =
      if remaining = 0 then
        Utils.has_list_in_list current_list completed_lists 0
      else -1
    in
    if remaining = 0 && found_dup_index > -1 then
      if is_part_2 then found_dup_index + 1 else cycle
    else
      match remaining with
      | 0 ->
          (* no more to distribute, add current list to completed lists and get the max again *)
          let max_index = get_max_pos current_list in
          let max_value = List.nth current_list max_index in
          let new_current =
            current_list
            |> List.mapi (fun i e -> if i = max_index then 0 else e)
          in
          redistribute (cycle + 1) max_value (get_next_index max_index)
            new_current
            (if cycle = 0 then completed_lists
             else current_list :: completed_lists)
      | _ ->
          (* update the current list *)
          let new_current =
            current_list
            |> List.mapi (fun i e -> if i = current_index then e + 1 else e)
          in
          redistribute cycle (remaining - 1)
            (get_next_index current_index)
            new_current completed_lists
  in
  redistribute 0 0 0 input []

let part1 = string_of_int @@ find_duplicate false
let part2 = string_of_int @@ find_duplicate true
