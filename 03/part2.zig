const std = @import("std");
const stdout = std.io.getStdOut().writer();
var gpallocator = std.heap.GeneralPurposeAllocator(.{}){};
var allocator = &gpallocator.allocator;

pub fn main() !void {
    const f = try std.fs.cwd().openFile("input", .{ .read = true });
    const b = try f.reader().readAllAlloc(allocator, 1024000);

    var tree_map_list = std.ArrayList([]bool).init(allocator);
    var trees_list = std.ArrayList(bool).init(allocator);

    for (b) |char| {
        if (char != '\n') {
            try trees_list.append(char == '#');
            continue;
        }

        var trees = trees_list.toOwnedSlice();
        try tree_map_list.append(trees);
        trees_list = std.ArrayList(bool).init(allocator);
    }

    var tree_map = tree_map_list.toOwnedSlice();

    var pairs = [_][2]usize{
        .{ 1, 1 },
        .{ 3, 1 },
        .{ 5, 1 },
        .{ 7, 1 },
        .{ 1, 2 },
    };
    var bumps: usize = 1;

    for (pairs) |pair| {
        bumps *= findBump(tree_map, pair[0], pair[1]);
    }

    try stdout.print("{}\n", .{bumps});
}

pub fn findBump(tree_map: [][]bool, right: usize, down: usize) usize {
    var right_cum = right;
    var down_cum = down;
    var bumps: usize = 0;

    while (down_cum < tree_map.len) {
        if (tree_map[down_cum][right_cum]) {
            bumps += 1;
        }

        right_cum = (right_cum + right) % tree_map[down].len;
        down_cum += down;
    }

    return bumps;
}
