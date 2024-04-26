import Slider from "rc-slider";
import Tooltip from "rc-tooltip";
import { Bar } from "react-chartjs-2";
import "chart.js/auto";
import "rc-slider/assets/index.css";
import "rc-tooltip/assets/bootstrap.css";
import _ from "lodash";
import { useMemo, useState } from "react";

const GRAPH_WIDTH = "768px";
const GRAPH_HEIGHT = "240px";

function Graph({ data }: { data: number[] }) {
  const marks: Record<number, React.ReactNode> = _.mapValues(
    _.range(11),
    () => <></>
  );
  const labels: string[] = useMemo(() => {
    const l: string[] = [];
    _.range(0, 10).forEach((i) => {
      _.range(0, 10).forEach((j) => {
        l.push(`${i}.${j}`);
      });
    });
    l.push("10.0");
    return l;
  }, []);
  const [handler, setHandler] = useState<number[]>([]);

  return (
    <>
      <div>
        <Bar
          width={GRAPH_WIDTH}
          height={GRAPH_HEIGHT}
          options={{
            responsive: true,
            scales: {
              x: {
                min: 0,
                max: 100,
                offset: true,
                grid: {
                  offset: false,
                  drawOnChartArea: false,
                },
                ticks: {
                  callback: function (_value, index) {
                    if (index % 10 === 0) {
                      return this.getLabelForValue(index);
                    }
                  },
                },
              },
              y: {
                min: 0,
                ticks: {
                  stepSize: 20,
                },
                grid: {
                  drawOnChartArea: false,
                },
              },
            },
            plugins: {
              legend: {
                display: false,
              },
            },
            animation: false,
          }}
          data={{
            labels: labels,
            datasets: [
              {
                data: data,
                backgroundColor: (ctx) => {
                  const v = ctx.dataIndex;
                  const min = Math.min(...handler) * 10;
                  const max = Math.max(...handler) * 10;
                  if (v >= min && v <= max) {
                    return "#5E81AC";
                  }
                  return "#D8DEE9";
                },
              },
            ],
          }}
        />
      </div>
      <div className="sm:ml-[37px] ml-[29px] mr-[14px]">
        {/* https://github.com/react-component/slider/issues/856 */}
        <Slider
          range
          marks={marks}
          step={0.1}
          min={0}
          max={10}
          handleRender={(node, props) => (
            <>
              <Tooltip
                overlayInnerStyle={{ minHeight: "auto" }}
                overlay={props.value}
                placement="top"
              >
                {node}
              </Tooltip>
            </>
          )}
          styles={{
            track: {
              backgroundColor: "#81A1C1",
            },
            rail: {
              backgroundColor: "rgb(198,212,227)",
            },
            handle: {
              backgroundColor: "#81A1C1",
              borderColor: "#81A1C1",
              opacity: 1,
            },
          }}
          dotStyle={{
            backgroundColor: "rgb(198,224,227)",
            borderColor: "rgb(198,224,227)",
          }}
          activeDotStyle={{
            backgroundColor: "#81A1C1",
            borderColor: "#81A1C1",
          }}
          allowCross={true}
          defaultValue={[0, 20]}
          onChange={(nums: number | number[]) => {
            setHandler(nums as number[]);
            console.log(nums);
          }}
        />
      </div>
    </>
  );
}

export default Graph;
