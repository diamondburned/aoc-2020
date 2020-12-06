const std = @import("std");
const stdout = std.io.getStdOut().writer();
var gpallocator = std.heap.GeneralPurposeAllocator(.{}){};
var allocator = &gpallocator.allocator;

// This code was written AFTER the Go iterations.

pub fn main() !void {
    const f = try std.fs.cwd().openFile("input", .{ .read = true });
    const b = try f.reader().readAllAlloc(allocator, 1024000);

    var group_iterator = std.mem.split(b, "\n\n");
    var total: usize = 0;

    while (true) {
        var group = group_iterator.next() orelse break;

        var answer = [_]u32{0} ** 256;
        var people: usize = 1;

        for (group) |char, i| {
            if (i != group.len - 1 and char == '\n') {
                people += 1;
                continue;
            }

            answer[char] += 1;
        }

        for (answer) |count| {
            if (count == people) {
                total += 1;
            }
        }
    }

    try stdout.print("{}\n", .{total});
}
