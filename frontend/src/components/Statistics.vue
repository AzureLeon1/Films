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
      countryCount: []
    };
  },
  methods: {
    getStatic() {
      this.allFilms.forEach(ele => {
        ele.genres.forEach(g => {
          if (this.ctg.indexOf(g) == -1) {
            this.ctg.push(g);
            this.ctgCount.push(1);
          } else {
            this.ctgCount[this.ctg.indexOf(g)]++;
          }
        });

        ele.countries.forEach(g => {
          if (this.country.indexOf(g) == -1) {
            this.country.push(g);
            this.countryCount.push(1);
          } else {
            this.countryCount[this.country.indexOf(g)]++;
          }
        });
      });
      console.log(this.country);
      console.log(this.countryCount);
    },
    getData: function() {
      this.$axios
        .get("/statistics", {
          methods: "get",
          baseURL: "http://www.leonwang.top:2333"
          // baseURL: 'http://www.leonwang.top:2333',
        })
        .then(result => {
          console.log(result);
          this.allFilms = result.data;
          this.getStatic();
          this.drawCategory();
          this.drawCountry();
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
      // 处理数据，方便只取出排序靠前的数据进行可视化
      var ctgObj = [];
      this.ctg.forEach(ele => {
        ctgObj.push({
          ctg: ele,
          ctgCount: this.ctgCount[this.ctg.indexOf(ele)]
        });
      });
      function compare(property) {
        return function(obj1, obj2) {
          var value1 = obj1[property];
          var value2 = obj2[property];
          return value2 - value1; // 降序
        };
      }
      var sortedCtgObj = ctgObj.sort(compare("ctgCount"));
      var sortedCtg = [];
      var sortedCtgCount = [];
      sortedCtgObj.forEach(ele => {
        sortedCtg.push(ele.ctg);
        sortedCtgCount.push(ele.ctgCount);
      });

      // 基于准备好的dom，初始化echarts实例
      let myChart = this.$echarts.init(
        document.getElementById("category"),
        "vintage"
      );
      // 绘制图表
      myChart.setOption({
        title: { text: "各类型影片数量（前25名）" },
        tooltip: {},
        xAxis: {
          type: "category",
          data: sortedCtg.slice(0, 25),
          axisLabel: {
            interval: 0,
            rotate: 40
          }
        },

        yAxis: {
          type: "value"
        },
        series: [
          {
            name: "影片数量",
            type: "bar",
            data: sortedCtgCount.slice(0, 25)
          }
        ]
      });
    },
    drawCountry() {
      // 处理数据，方便只取出排序靠前的数据进行可视化
      var countryObj = [];
      this.country.forEach(ele => {
        countryObj.push({
          country: ele,
          countryCount: this.countryCount[this.country.indexOf(ele)]
        });
      });
      function compare(property) {
        return function(obj1, obj2) {
          var value1 = obj1[property];
          var value2 = obj2[property];
          return value2 - value1; // 降序
        };
      }
      var sortedCountryObj = countryObj.sort(compare("countryCount"));
      var sortedCountry = [];
      var sortedCountryCount = [];
      sortedCountryObj.forEach(ele => {
        sortedCountry.push(ele.country);
        sortedCountryCount.push(ele.countryCount);
      });


      // 基于准备好的dom，初始化echarts实例
      let myChart = this.$echarts.init(
        document.getElementById("country"),
        "dark"
      );
      // 绘制图表
      myChart.setOption({
        title: { text: "各国家/地区影片数量（前30名）" },
        tooltip: {},
        xAxis: {
          type: "category",
          data: sortedCountry.slice(0,30),
          axisLabel: {
            interval: 0,
            rotate: 40
          }
        },
        yAxis: {
          type: "value"
        },
        series: [
          {
            name: "影片数量",
            type: "bar",
            data: sortedCountryCount.slice(0,30)
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
