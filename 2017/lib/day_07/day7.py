from collections import defaultdict, Counter

def find_root(children):
    # All nodes in the tree minus all nodes that are children.
    # N.B. sum(list of lists, []) flattens a list of lists because of how + is
    # overloaded in Python.
    root = set(children.keys()) - set(sum(children.values(),[]))
    # The root should be a single item, so we should be able to safely pop it.
    return root.pop()

def find_weight(weights, children, node, memo={}):
    # Recursively add the weight of the node to the weight of its children.
    return weights[node] + sum([find_weight(weights, children, child, memo) for
                                child in children[node]])

def find_imbalance(weights, children, node):
    # If we're at a leaf node, we can't find an imbalance.
    if not children[node]: return 
    child_weights = [find_weight(weights, children, child) for child in
                     children[node]]
    weight_counts = Counter(child_weights)
    # Exit early if everything is the same.
    if len(weight_counts) <= 1: return
    # The least frequent value will be the differing value.
    # N.B. Counter.most_common() returns all counts, ordered in reverse by
    # frequency. [-1] gets the last, and [0] gets the weight.
    least_frequent_weight = weight_counts.most_common()[-1][0]
    # N.B. child_weights and children[node] are in the same order.
    lfw_node = children[node][child_weights.index(least_frequent_weight)]
    # Try to find an imbalance deeper into the tree (we want to get to the leaf
    # of the problem).
    deeper_imbalance = find_imbalance(weights, children, lfw_node)
    if deeper_imbalance:
        return deeper_imbalance
    else:
        most_frequent_weight = weight_counts.most_common(1)[0][0]
        imbalance = most_frequent_weight - least_frequent_weight
        return weights[lfw_node] + imbalance


with open('input.txt') as inp:
    # The weight of a node.
    weights = {}
    # Adjacency list with node as the key.
    children = defaultdict(list)
    for line in inp.readlines():
        # Lines look like 'node (value) -> child1 child2 child3 ...'
        node, weight, *child_nodes = line.strip().split(' ')
        # Strip parentheses and convert an int.
        weight = int(weight.strip('()'))
        weights[node] = weight
        if child_nodes:
            # Remove the arrow ('->') and the commas after each child.
            children[node] = [child.strip(',') for child in child_nodes[1:]]

    # Part 1.
    root = find_root(children)
    print(root)

    # Part 2.
    print(find_imbalance(weights, children, root))

