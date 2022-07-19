from pyecharts.charts import Line
import pyecharts.options as opts


def draw_line(title, xaxis, yaxis):
    line = Line(init_opts=opts.InitOpts(width="1600px", height="800px"))
    line.add_xaxis(xaxis_data=xaxis)
    line.add_yaxis(y_axis=yaxis, series_name="speed", is_smooth=True)
    line.set_global_opts(title_opts=opts.TitleOpts(title=title))
    return line
