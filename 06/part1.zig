const std = @import("std");
const stdout = std.io.getStdOut().writer();
var gpallocator = std.heap.GeneralPurposeAllocator(.{}){};
var allocator = &gpallocator.allocator;

// This code was written AFTER the Go iterations.

fn arrayLen(comptime T: type, array: []const T, falsy: T) usize {
    var pass: usize = 0;

    for (array) |v| {
        if (v != falsy) {
            pass += 1;
        }
    }

    return pass;
}

pub fn main() !void {
    const f = try std.fs.cwd().openFile("input", .{ .read = true });
    const b = try f.reader().readAllAlloc(allocator, 1024000);

    var group_iterator = std.mem.split(b, "\n\n");
    var total: usize = 0;

    while (true) {
        var group = group_iterator.next() orelse break;
        var answer = [_]bool{false} ** 256;

        for (group) |char| {
            if (char != '\n') {
                answer[char] = true;
            }
        }

        total += arrayLen(bool, &answer, false);
    }

    try stdout.print("{}\n", .{total});
}
