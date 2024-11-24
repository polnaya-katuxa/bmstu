import plotly.graph_objects as go

with open('out/data.txt') as f:
    x = [float(e) for e in f.readline().split()]
    y = [float(e) for e in f.readline().split()]
    z = [[float(e) for e in l.split()] for l in f.readlines()]

    fig = go.Figure(data=[go.Surface(x=x,
                                    y=y,
                                    z=z,
                                    )])

    margin = dict(r=20, l=20, b=20, t=20)
    fig.update_layout(margin=margin, scene_aspectmode='manual', scene_aspectratio=dict(x=2, y=2, z=2), scene = dict(zaxis = dict(range=[290,320],),))

    fig.show()