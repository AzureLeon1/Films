<template>
  <div>
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
                v-model="film.rating.average"
                disabled
                show-score
                text-color="#ff9900"
                score-template="{value}"
              ></el-rate>
            </div>
            <div class="text item">{{'国家：' + getAllElements(film.countries)}}</div>
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
      :total="this.allFilms.length"
      :page-size="10"
      :current-page.sync="currentPage"
      style="margin-top: 10px"
    ></el-pagination>
  </div>
</template>

<script>
export default {
  name: "FilmList",
  data() {
    return {
      allFilms: [],
      currentFilms: [],
      currentPage: 1
    };
  },
  methods: {
    getData: function() {
      this.$axios
        .get("../../static/films.json")
        .then(response => {
          response.data.split("\n").forEach(element => {
            this.allFilms.push(JSON.parse(element));
            // 电影评分数据类型由String转换没Number，由于评分插件需要
            this.allFilms[this.allFilms.length - 1].rating.average = Number(
              this.allFilms[this.allFilms.length - 1].rating.average
            );
          });
          this.currentFilms = this.allFilms.slice(0, 10);
        })
        .catch(err => {
          console.log(err);
        });
    },
    handleCurrentChange() {
      this.currentFilms = this.allFilms.slice(
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
    }
  },
  mounted() {
    this.allFilms = [];
    this.getData();
    console.log(this.allFilms);
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.whole {
  text-align: center;
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
  width: 900px;
  margin: 15px;
  margin-left: 25%;
}
</style>
