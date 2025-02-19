module StringMap = Map.Make (String)

type program = {
  id : string;
  weight: int;
  children: string list;
}

type program2 = {
  id : string;
  parent_id : string;
  weight: int;
  depth: int;
}

let input : program list = 
  Utils.read_whole_file "./lib/day_07/input.txt" 
  |> String.split_on_char '\n'
  |> List.map (fun row -> 
      let regex = Str.regexp "->" in
      let id_child_pair = Str.split regex row in
      let id_weight_pair = id_child_pair |> List.hd |> String.split_on_char ' ' in
      let id = id_weight_pair |> List.hd in
      let weight = id_weight_pair 
        |> List.tl 
        |> List.hd 
        |> Str.replace_first (Str.regexp "(") "" 
        |> Str.replace_first (Str.regexp ")") "" 
        |> int_of_string in

      if List.length id_child_pair > 1 then
        let children = List.nth id_child_pair 1 
          |> String.split_on_char ',' 
          |> List.map (fun c -> c 
          |> String.trim) in
        { id = id; weight = weight; children = children; }
      else (* no child *)
        { id = id; weight = weight; children = []; }
      )

  

let part1 =
  let rec find_parent id =
    let parent = input |> List.find_opt (fun p -> p.children |> List.exists (fun c -> c = id)) in
    match parent with
    | None -> id
    | Some p -> find_parent p.id
  in
  let first_edge = input |> List.filter (fun r -> r.children = [] ) |> List.hd in
  let parent = find_parent first_edge.id in
  if parent = first_edge.id then "NA" else parent

let part2 = "N/A"
