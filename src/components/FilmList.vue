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
            <el-menu-item index="1">电影列表</el-menu-item>
            <el-menu-item index="3">分类</el-menu-item>
          </el-menu>

          <!-- <div style="margin: 5px 20px 0">
      <el-button type="primary" size="small" plain @click="changeStandard(1)">按名称</el-button>
      <el-button type="success" size="small" plain @click="changeStandard(2)">按导演</el-button>
      <el-button type="info" size="small" plain @click="changeStandard(3)">按演员</el-button>
      <el-button type="warning" size="small" plain @click="changeStandard(4)">按剧情</el-button>
          </div>-->
        </div>
      </el-header>

      <el-container>
        <el-aside width="360px">
          <div style="margin-top: 20px; margin-left: 15px;">
            <el-card class="rank-box">
              <div class="clearfix">
                <span class="title">排行榜</span>
              </div>
              <div>
                <el-table
                  :data="allFilms"
                  style="width: 100%"
									height="2000"
                  :default-sort="{prop: 'aveRating', order: 'descending'}"
                >
									<el-table-column prop="title" label="名称" width="225"></el-table-column>
                  <el-table-column prop="aveRating" label="评分" sortable width="75"></el-table-column>
                  
                </el-table>
              </div>
            </el-card>
          </div>
        </el-aside>
        <el-main>
          <div style="margin-top: 5px; margin-left: 15px;">
            <el-input class="search" placeholder="搜索电影" v-model="filterText"></el-input>
            <el-radio-group v-model="defaultStandard" size="medium" @change="changeStandard">
              <el-radio-button label="按名称"></el-radio-button>
              <el-radio-button label="按导演"></el-radio-button>
              <el-radio-button label="按演员"></el-radio-button>
              <el-radio-button label="按剧情"></el-radio-button>
            </el-radio-group>
          </div>
          <div v-for="film in currentFilms" :key="film._id">
            <el-card class="box-card" :body-style="{ padding: '10px' }">
              <div slot="header" class="clearfix">
                <span class="title">{{film.title}}</span>
              </div>
              <el-container>
                <el-aside width="200px">
                  <img
                    :src="film.poster"
                    class="image"
                    onerror="this.src='https://ws3.sinaimg.cn/large/006tNc79ly1g2bpa1n66mj3050050glo.jpg'"
                  >
                </el-aside>
                <el-main>
                  <div class="text item">
                    <el-rate
                      v-model="film.aveRating"
                      disabled
                      show-score
                      text-color="#ff9900"
                      score-template="{value}"
                    ></el-rate>
                  </div>
                  <div class="text item">{{getAllElements(film.genres) + ' （'+ getAllElements(film.countries) + '）'}}</div>
                  <div class="text item">{{'上映时间：' + film.pubdate}}</div>
                  <div class="text item">{{'导演：' + getAllNames(film.directors)}}</div>
                  <div class="text item name">{{'主演：' + getAllNames(film.casts)}}</div>
                  <div class="text item">{{'剧情：' + film.summary}}</div>
                  <el-button style="float: right; padding: 3px 0" type="text">查看详情</el-button>
                </el-main>
              </el-container>
            </el-card>
          </div>

          <!-- 每行两个卡片的布局方式，不太美观 -->
          <!-- <el-row v-for="o1 in 5" :key="o1" >
        <el-col :span="8" v-for="(o2, index2) in 2" :key="o2" :offset="index2 > 0 ? 2 : 0">
          <el-card :body-style="{ padding: '0px' }">
            <img :src="currentFilms[o1*2+o2-3].poster" class="image" onerror="this.src='https://ws3.sinaimg.cn/large/006tNc79ly1g2bpa1n66mj3050050glo.jpg'">
            <div style="padding: 14px;">
              <span>{{ currentFilms[o1*2+o2-3].title }}</span>
              <div class="bottom clearfix">
                <time class="time">{{currentFilms[o1*2+o2-3].pubdate}} </time>
                <el-button type="text" class="button">查看详情</el-button>
              </div>
            </div>
          </el-card>
        </el-col>
          </el-row>-->

          <el-pagination
            background
            @current-change="handleCurrentChange"
            layout="total, prev, pager, next"
            :total="this.filterFilms.length"
            :page-size="10"
            :current-page.sync="currentPage"
            style="margin-top: 10px"
          ></el-pagination>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
export default {
  name: "FilmList",
  data() {
    return {
      activeIndex: "1",
      defaultStandard: "按名称",
      selectedStandard: "按名称",
      filterText: "",
      allFilms: [],
      filterFilms: [],
      currentFilms: [],
      currentPage: 1
    };
  },
  watch: {
    // 过滤关键字发生变化时，自动过滤一遍
    filterText(val) {
      console.log(this.selectedStandard);
      if (this.selectedStandard == "按名称") {
        this.filterFilms = this.allFilms.filter(ele => {
          return ele.title.indexOf(val) != -1;
        });
      }
      if (this.selectedStandard == "按导演") {
        this.filterFilms = this.allFilms.filter(ele => {
          return this.getAllNames(ele.directors).indexOf(val) != -1;
        });
      }
      if (this.selectedStandard == "按演员") {
        this.filterFilms = this.allFilms.filter(ele => {
          return this.getAllNames(ele.casts).indexOf(val) != -1;
        });
      }
      if (this.selectedStandard == "按剧情") {
        this.filterFilms = this.allFilms.filter(ele => {
          return ele.summary.indexOf(val) != -1;
        });
      }
      console.log(this.selectedStandard);
      this.handleCurrentChange();
      console.log(this.filterFilms);
    },
    // 过滤标准发生变化时，根据原关键字按照新的过滤标准重新过滤一遍
    selectedStandard(val) {
      if (val == "按名称") {
        this.filterFilms = this.allFilms.filter(ele => {
          return ele.title.indexOf(this.filterText) != -1;
        });
      }
      if (val == "按导演") {
        this.filterFilms = this.allFilms.filter(ele => {
          return this.getAllNames(ele.directors).indexOf(this.filterText) != -1;
        });
      }
      if (val == "按演员") {
        this.filterFilms = this.allFilms.filter(ele => {
          return this.getAllNames(ele.casts).indexOf(this.filterText) != -1;
        });
      }
      if (val == "按剧情") {
        this.filterFilms = this.allFilms.filter(ele => {
          return ele.summary.indexOf(this.filterText) != -1;
        });
      }
      console.log(this.selectedStandard);
      this.handleCurrentChange();
      console.log(this.filterFilms);
    }
  },
  methods: {
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
          });
          this.filterFilms = this.allFilms;
          this.currentFilms = this.filterFilms.slice(0, 10);
        })
        .catch(err => {
          console.log(err);
        });
    },
    handleCurrentChange() {
      this.currentFilms = this.filterFilms.slice(
        (this.currentPage - 1) * 10,
        (this.currentPage - 1) * 10 + 10
      );
    },
    getAllNames(group) {
      let names = "";
      group.forEach(ele => {
        names += ele.name;
        if (group.indexOf(ele) != group.length - 1) {
          names += " / ";
        }
      });
      return names;
    },
    getAllElements(array) {
      let all = "";
      array.forEach(ele => {
        all += ele;
        if (array.indexOf(ele) != array.length - 1) {
          all += " / ";
        }
      });
      return all;
    },
    changeStandard(val) {
      console.log(val);
      this.selectedStandard = val;
    },
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    }
  },
  created() {},
  mounted() {
    this.allFilms = [];
    this.filterFilms = [];
    this.currentFilms = [];
    this.filterText = "";
    this.getData();
    console.log(this.allFilms);
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.nav {
  width: 100%;
}

.search {
  width: 450px;
  margin-bottom: 5px;
}

.image {
  margin-left: 30px;
  margin-top: 25px;
  width: 60%;
  display: block;
}

.title {
  color: #0984e3;
  font-weight: 500;
  font-size: 20px;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}

.clearfix:after {
  clear: both;
}
.text {
  font-size: 14px;
  text-align: left;
}

.item {
  margin-bottom: 5px;
  overflow: hidden;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.name {
  -webkit-line-clamp: 1;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}
.clearfix:after {
  clear: both;
}

.box-card {
  text-align: center;
  margin: 15px;
}

.rank-box {
  height: 100%;
}
</style>
