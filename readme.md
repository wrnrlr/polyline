# Polyline

A module for [Gio](https://gioui.org) to draw polylines.


It's currently a very naive implementation that
draws every line segment and joins them with a circle.
This solution is not suitable for transparent colors because
overlapping parts are painted repeatedly resulting in an uneven coloring. 

Hopefully we can fix this in the future and keep using the same api.

## API

```
line := []f32.Point{{20, 20}, {20, 70}, {80, 80}, {110, 100}, {300, 500}}
polyline.Draw(line, width, color, gtx)
```

See `example` on how to use this to make a paint app.