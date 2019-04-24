<template>
  <div>
    <el-container>
      <el-header style="padding: 0;">
        <div class="nav">
          <el-menu
            :default-active="activeIndex"
            class="el-menu-demo"
            mode="horizontal"
            @select="handleSelect"
            background-color="#545c64"
            text-color="#fff"
            active-text-color="#ffd04b"
          >
            <el-menu-item index="1" style="font-size: 20px">电影列表</el-menu-item>
            <el-menu-item index="2" style="font-size: 20px">分类</el-menu-item>
            <el-menu-item index="3" style="font-size: 20px">数据统计</el-menu-item>
          </el-menu>
        </div>
      </el-header>
      <el-main>
        <div class="main-part">
          <el-card class="box-card" :body-style="{ padding: '10px' }">
            <div style="margin: 20px 100px;">
              <div class="text item">
                <i class="el-icon-edit" style="margin: 15px;"></i>
                {{'影片数量：' + allFilms.length + " 部"}}
              </div>

              <div id="category" :style="{width: '800px', height: '500px'}"></div>
              <div id="country" :style="{width: '800px', height: '500px'}"></div>
            </div>
          </el-card>
        </div>
      </el-main>
      <el-main></el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  name: "Statistics",
  data() {
    return {
      activeIndex: "3",
      allFilms: [],
      ctg: [],
      ctgCount: [],
      country: [],
      countryCount: [],
    };
  },
  methods: {
    getStatic() {
      this.allFilms.forEach(ele => {
        ele.genres.forEach(g => {
          if(this.ctg.indexOf(g) == -1) {
            this.ctg.push(g)
            this.ctgCount.push(1)
          }
          else {
            this.ctgCount[this.ctg.indexOf(g)]++
          }
        })

        ele.countries.forEach(g => {
          if(this.country.indexOf(g) == -1) {
            this.country.push(g)
            this.countryCount.push(1)
          }
          else {
            this.countryCount[this.country.indexOf(g)]++
          }
        })
      })
      console.log(this.country);
      console.log(this.countryCount);
    },
    getData: function() {
      this.$axios
        .get("../../static/films.json")
        .then(response => {
          response.data.split("\n").forEach(element => {
            this.allFilms.push(JSON.parse(element));
            // 电影评分数据类型由String转换没Number，由于评分插件需要
            this.allFilms[this.allFilms.length - 1].aveRating = Number(
              this.allFilms[this.allFilms.length - 1].rating.average
            );
            this.allFilms[this.allFilms.length - 1].stars =
              this.allFilms[this.allFilms.length - 1].aveRating / 2;
          });
          return this.allFilms
        })
        .then( () => {
          // 如果直接写在mounted函数中会由于axios的异步，getStatic()无法分析出数据
          this.getStatic()
        })
        .then( () => {
          this.drawCategory()
          this.drawCountry()
        })
        .catch(err => {
          console.log(err);
        });
    },
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
      if (keyPath[0] == "1" || "2") {
        this.$router.push({
          name: "list",
          params: {
            index: keyPath[0]
          }
        });
      }
    },
        drawCategory() {
      // 基于准备好的dom，初始化echarts实例
      let myChart = this.$echarts.init(document.getElementById("category"),'vintage');
      // 绘制图表
      myChart.setOption({
        title: { text: "各类型影片数量" },
        tooltip: {},
        xAxis: {
          type: 'category',
          data: this.ctg
        },
        yAxis: {
          type: 'value'
        },
        series: [
          {
            name: "影片数量",
            type: "bar",
            data: this.ctgCount
          }
        ]
      });
    },
     drawCountry() {
      // 基于准备好的dom，初始化echarts实例
      let myChart = this.$echarts.init(document.getElementById("country"),'dark');
      // 绘制图表
      myChart.setOption({
        title: { text: "各国家/地区影片数量" },
        tooltip: {},
        xAxis: {
          type: 'category',
          data: this.country
        },
        yAxis: {
          type: 'value'
        },
        series: [
          {
            name: "影片数量",
            type: "bar",
            data: this.countryCount
          }
        ]
      });
    }
  },
  mounted() {
    this.allFilms = [];
    this.getData();
    // console.log(this.allFilms);
    // this.drawLine();
  }
};
</script>

<style scoped>
.nav {
  width: 100%;
}

.box-card {
  text-align: center;
  margin: 15px 150px;
}

.text {
  font-size: 20px;
  text-align: left;
}

.item {
  margin-bottom: 5px;
  overflow: hidden;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}
</style>
