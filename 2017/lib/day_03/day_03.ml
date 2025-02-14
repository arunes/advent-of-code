type direction =
  | East
  | North
  | West
  | South

type point = {
  x : int;
  y : int;
  value: int;
}

let part_2 target =
  let get_value x y matrix = 
    let n_n = 
      match List.find_opt (fun p -> p.x = x && p.y = (y+1)) matrix with
      | Some p ->  p.value
      | None ->  0
    in

    let n_ne = 
      match List.find_opt (fun p -> p.x = (x+1) && p.y = (y+1)) matrix with
      | Some p ->  p.value
      | None -> 0
    in

    let n_e = 
      match List.find_opt (fun p -> p.x = (x+1) && p.y = y) matrix with
      | Some p ->  p.value
      | None -> 0
    in

    let n_se = 
      match List.find_opt (fun p -> p.x = (x+1) && p.y = (y-1)) matrix with
      | Some p ->  p.value
      | None -> 0
    in

    let n_s = 
      match List.find_opt (fun p -> p.x = x && p.y = (y-1)) matrix with
      | Some p ->  p.value
      | None -> 0
    in
    
    let n_sw = 
      match List.find_opt (fun p -> p.x = (x-1) && p.y = (y-1)) matrix with
      | Some p ->  p.value
      | None -> 0
    in

    let n_w = 
      match List.find_opt (fun p -> p.x = (x-1) && p.y = y) matrix with
      | Some p ->  p.value
      | None -> 0
    in

    let n_nw = 
      match List.find_opt (fun p -> p.x = (x-1) && p.y = (y+1)) matrix with
      | Some p ->  p.value
      | None -> 0
    in

    n_n + n_ne + n_e + n_se + n_s + n_sw + n_w + n_nw
  in

  let rec fill x y dir matrix = 
    let new_point = match dir with
      | East -> {x=x+1; y=y; value=get_value (x+1) y matrix}
      | North -> {x=x; y=y+1; value=get_value x (y+1) matrix}
      | West -> {x=x-1; y=y; value=get_value (x-1) y matrix}
      | South -> {x=x; y=y-1; value=get_value x (y-1) matrix}
    in 

    let new_dir = match dir with
      | East -> 
          if List.exists (fun p -> p.x=new_point.x && p.y = new_point.y + 1) matrix then East else North
      | North ->
          if List.exists (fun p -> p.x=new_point.x - 1 && p.y = new_point.y) matrix then North else West
      | West ->
          if List.exists (fun p -> p.x=new_point.x && p.y = new_point.y - 1) matrix then West else South
      | South -> 
          if List.exists (fun p -> p.x=new_point.x + 1 && p.y = new_point.y) matrix then South else East
    in   

    if new_point.value >= target then
      new_point.value
    else 
      fill new_point.x new_point.y new_dir (new_point :: matrix)
    in

  fill 0 0 East [{ x= 0; y=0; value=1 }]


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
let part1 = string_of_int @@ find_distance input
let part2 = string_of_int @@ part_2 input
