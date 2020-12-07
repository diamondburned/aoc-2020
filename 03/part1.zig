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

    var right: usize = 3;
    var down: usize = 1;
    var bump: usize = 0;

    while (down < tree_map.len) {
        if (tree_map[down][right]) {
            bump += 1;
        }

        right = (right + 3) % tree_map[down].len;
        down += 1;
    }

    try stdout.print("{}\n", .{bump});
}
