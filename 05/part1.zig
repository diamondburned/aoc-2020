const std = @import("std");
const stdout = std.io.getStdOut().writer();
var gpallocator = std.heap.GeneralPurposeAllocator(.{}){};
var allocator = &gpallocator.allocator;

// This code was written AFTER the Go iterations.

pub fn bseat(input: [10]u8) ![2]u16 {
    var binary_buf = [_]u8{0} ** 10;

    for (input) |char, index| {
        switch (char) {
            'F', 'L' => binary_buf[index] = '0',
            'B', 'R' => binary_buf[index] = '1',
            else => unreachable,
        }
    }

    return [2]u16{
        try std.fmt.parseInt(u16, binary_buf[0..7], 2),
        try std.fmt.parseInt(u16, binary_buf[7..10], 2),
    };
}

pub fn seatId(coord: [2]u16) u16 {
    return coord[0] * 8 + coord[1];
}

pub fn main() !void {
    const f = try std.fs.cwd().openFile("input", .{ .read = true });
    const b = try f.reader().readAllAlloc(allocator, 1024000);

    var max_seat: u16 = 0;
    var seat_cur: usize = 0;
    var seat_buf: [10]u8 = [_]u8{0} ** 10;

    for (b) |char| {
        if (char != '\n') {
            seat_buf[seat_cur] = char;
            seat_cur += 1;
            continue;
        }

        seat_cur = 0;

        var coord = try bseat(seat_buf);
        var seat_id = seatId(coord);

        if (seat_id > max_seat) {
            max_seat = seat_id;
        }
    }

    try stdout.print("{}\n", .{max_seat});
}
