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
          <div v-if="activeIndex == 1" style="margin-top: 5px; margin-left: 15px;">
            <el-input class="search" placeholder="搜索电影" v-model="filterText"></el-input>
            <el-button round size="medium" @click="search">搜索</el-button>
            <div>
              请选择搜索标准：
              <el-radio-group v-model="defaultStandard" size="medium" @change="changeStandard">
                <el-radio-button label="按名称"></el-radio-button>
                <el-radio-button label="按导演"></el-radio-button>
                <el-radio-button label="按演员"></el-radio-button>
                <el-radio-button label="按剧情"></el-radio-button>
              </el-radio-group>
            </div>
          </div>
          <div v-if="activeIndex == 2" style="margin-top: 5px; margin-left: 15px;">
            <el-radio-group v-model="defaultCategory" size="medium" @change="changeCategory">
              <el-radio-button label="全部"></el-radio-button>
              <el-radio-button label="喜剧"></el-radio-button>
              <el-radio-button label="科幻"></el-radio-button>
              <el-radio-button label="动作"></el-radio-button>
              <el-radio-button label="犯罪"></el-radio-button>
              <el-radio-button label="惊悚"></el-radio-button>
              <el-radio-button label="爱情"></el-radio-button>
              <el-radio-button label="传记"></el-radio-button>
              <el-radio-button label="纪录片"></el-radio-button>
            </el-radio-group>
          </div>
          <div v-for="film in currentFilms" :key="film._id">
            <el-card class="box-card" :body-style="{ padding: '0px'}">
              <!-- <div slot="header" class="clearfix">
                <span class="title" @click="viewDetails(film)">{{film.title}}</span>
              </div>-->
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
                    <el-row class="title">{{film.title}}</el-row>
                    <el-row>
                      <div style="width: 130px">
                        <el-col>
                          <div>
                            <el-rate
                              allow-half
                              v-model="film.stars"
                              disabled
                              text-color="#ff9900"
                              score-template="{value}"
                            ></el-rate>
                          </div>
                        </el-col>
                      </div>
                      <el-col :span="6">
                        <div style="color: #ff9900; font-size: 16px">{{film.aveRating}}</div>
                      </el-col>
                    </el-row>
                  </div>
                  <div
                    class="text item"
                  >{{getAllElements(film.genres) + ' （'+ getAllElements(film.countries) + '）'}}</div>
                  <div class="text item">{{'上映时间：' + film.pubdate}}</div>
                  <div class="text item">{{'导演：' + getAllNames(film.directors)}}</div>
                  <div class="text item name">{{'主演：' + getAllNames(film.casts)}}</div>
                  <div class="text item">{{'剧情：' + film.summary}}</div>
                  <el-button
                    style="float: right; padding: 3px 0"
                    type="text"
                    @click="viewDetails(film)"
                  >查看详情</el-button>
                </el-main>
              </el-container>
            </el-card>
          </div>

          <el-pagination
            background
            @current-change="handleCurrentChange"
            layout="total, prev, pager, next"
            :total="10000"
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
      defaultCategory: "全部",
      selectedCategory: "全部",
      defaultStandard: "按名称",
      selectedStandard: "按名称",
      filterText: "",
      allFilms: [],
      filterFilms: [],
      currentFilms: [],
      currentPage: 1
    };
  },
  methods: {
    getData: function() {
      if (this.$route.params.index != null) {
        this.activeIndex = this.$route.params.index;
      }

      this.$axios
        .get("/film", {
          methods: "get",
          baseURL: "http://www.leonwang.top:2333"
          // baseURL: 'http://www.leonwang.top:2333',
        })
        .then(result => {
          console.log(result);
          this.allFilms = result.data.filmsForRank;
          this.allFilms.forEach(ele => {
            ele.aveRating = Number(ele.rating.average);
            ele.stars = ele.aveRating / 2;
          });
          this.filterFilms = this.allFilms;
          this.currentFilms = result.data.films;
          this.currentFilms.forEach(ele => {
            ele.aveRating = Number(ele.rating.average);
            ele.stars = ele.aveRating / 2;
          });
        })
        .catch(err => {
          console.log(err);
        });

    },
    search() {
      var standard;
      if (this.selectedStandard == "按名称") {
        standard = "title";
      }
      if (this.selectedStandard == "按导演") {
        standard = "directors";
      }
      if (this.selectedStandard == "按演员") {
        standard = "casts";
      }
      if (this.selectedStandard == "按剧情") {
        standard = "summary";
        console.log("search:" + this.selectedStandard + this.filterText);
      }
      this.$axios
        .get(`/search?standard=${standard}&keyWord=${this.filterText}`, {
          methods: "get",
          baseURL: "http://www.leonwang.top:2333"
        })
        .then(result => {
          this.currentFilms = result.data;
          console.log(result);
        })
        .catch(err => {
          console.log(err);
        });
    },
    handleCurrentChange() {
      this.$axios
        .get(`/film/${this.currentPage}`, {
          methods: "get",
          baseURL: "http://www.leonwang.top:2333"
          // baseURL: 'http://www.leonwang.top:2333',
        })
        .then(result => {
          this.currentFilms = result.data;
          console.log(result);
        })
        .catch(err => {
          console.log(err);
        });
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
    changeCategory(val) {
      console.log(val);
      if (val == "全部") {
        val = "";
      }
      this.selectedCategory = val;
      this.$axios
        .get("/film?limit=200", {
          methods: "get",
          baseURL: "http://www.leonwang.top:2333"
        })
        .then(result => {
          console.log(result);
          this.filterFilms = this.allFilms;
          this.currentFilms = result.data.films;
          this.currentFilms.forEach(ele => {
            ele.aveRating = Number(ele.rating.average);
            ele.stars = ele.aveRating / 2;
          });

          this.filterFilms = this.currentFilms.filter(ele => {
            return this.getAllElements(ele.genres).indexOf(val) != -1;
          });
          this.currentFilms = this.filterFilms;
        })
        .catch(err => {
          console.log(err);
        });

      // this.handleCurrentChange();
    },
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
      if (keyPath[0] == "3") {
        // this.$router.push('/statistics')
        this.$router.push({
          name: "statistics",
          params: {}
        });
      }
      this.activeIndex = keyPath[0];
      this.filterFilms = this.allFilms;
      this.currentPage = 1;
      this.handleCurrentChange();
    },
    viewDetails(f) {
      this.$router.push({
        name: "detail",
        params: {
          film: f
        }
      });
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
  cursor: pointer;
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
  -webkit-line-clamp: 1;
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
  margin: 8px;
}

.rank-box {
  height: 100%;
}
</style>
