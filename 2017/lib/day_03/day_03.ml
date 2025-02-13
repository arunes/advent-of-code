let find_distance target = 
  let rec get_distance distance min count = 
    let max = min + count - 1 in
    if target > max then get_distance (distance + 1) (max + 1) (count + 8) else 
      let total_rows = distance * 2 in
      
      let right_min = min in
      let right_max = right_min + total_rows - 2 in (* removing corners *)
      let top_min = right_max + 2 in
      let top_max = top_min + total_rows - 2 in (* removing corners *)
      let left_min = top_max + 2 in
      let left_max = left_min + total_rows - 2 in (* removing corners *)
      let bottom_min = left_max + 2 in
      let bottom_max = bottom_min + total_rows - 2 in (* removing corners *)

      let find_distance m = 
          Int.abs ((distance - 1) - (m - target)) + distance in

      match target with
      | t when t >= right_min && t <= right_max -> 
          find_distance right_max
      | t when t >= top_min && t <= top_max -> 
          find_distance top_max
      | t when t >= left_min && t <= left_max -> 
          find_distance left_max
      | t when t >= bottom_min && t <= bottom_max -> 
          find_distance bottom_max
      | _ -> total_rows

    in

    (* initial distance 1, first number 2, and total amount of numbers starts at 8*)
    match target with
    | t when t > 1 -> get_distance 1 2 8
    | _ -> 0

let input = 312051
let part1 = string_of_int (find_distance input)
let part2 = "N/A"
