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
            <el-menu-item index="1" style="font-size: 22px">电影列表</el-menu-item>
            <el-menu-item index="2" style="font-size: 22px">分类</el-menu-item>
            <el-menu-item index="3" style="font-size: 22px">电影详情</el-menu-item>
          </el-menu>
        </div>
      </el-header>
			<el-main>
				<div class="main-part">
					<el-card class="box-card" :body-style="{ padding: '10px' }">
						<div slot="header" class="clearfix">
                <span class="title">{{this.currentFilm.title}}</span>
              </div>
              <el-container>
                <el-aside width="300px">
                  <img
                    :src="this.currentFilm.poster"
                    class="image"
                    onerror="this.src='https://ws3.sinaimg.cn/large/006tNc79ly1g2bpa1n66mj3050050glo.jpg'"
                  >
									<div class="stars">
										<el-row>
											<el-col :span="16">
                    <el-rate
											allow-half
                      v-model="currentFilm.stars"
                      disabled
                      text-color="#ff9900"
                      score-template="{value}"
                    ></el-rate>
										</el-col>
										<el-col :span="6">
											<div style="color: #ff9900; font-size: 22px">{{currentFilm.aveRating + "分"}}
											</div>
											</el-col>
										</el-row>
										<el-row>
											<div style="text-align:right; padding-right:50px;">
												{{"by " + currentFilm.rating.rating_people + " 人评价"}}
											</div>
										</el-row>
										
										<div>

											<el-row>
											<el-col :span="16">
                    <el-rate
											allow-half
                      v-model="v5"
                      disabled
                      text-color="#ff9900"
                      score-template="{value}"
											:colors="['#99A9BF', '#F7BA2A', '#FF9900']"
                    ></el-rate>
										</el-col>
										<el-col :span="1">
											<div style="color: #ff9900; font-size: 16px">{{this.stars5 + "%"}}
											</div>
											</el-col>
										</el-row>

										<el-row>
											<el-col :span="16">
                    <el-rate
											allow-half
                      v-model="v4"
                      disabled
                      text-color="#ff9900"
                      score-template="{value}"
											:colors="['#99A9BF', '#F7BA2A', '#FF9900']"
                    ></el-rate>
										</el-col>
										<el-col :span="1">
											<div style="color: #ff9900; font-size: 16px">{{this.stars4 + "%"}}
											</div>
											</el-col>
										</el-row>

										<el-row>
											<el-col :span="16">
                    <el-rate
											allow-half
                      v-model="v3"
                      disabled
                      text-color="#ff9900"
                      score-template="{value}"
											:colors="['#99A9BF', '#F7BA2A', '#FF9900']"
                    ></el-rate>
										</el-col>
										<el-col :span="1">
											<div style="color: #ff9900; font-size: 16px">{{this.stars3 + "%"}}
											</div>
											</el-col>
										</el-row>

										<el-row>
											<el-col :span="16">
                    <el-rate
											allow-half
                      v-model="v2"
                      disabled
                      text-color="#ff9900"
                      score-template="{value}"
											:colors="['#99A9BF', '#F7BA2A', '#FF9900']"
                    ></el-rate>
										</el-col>
										<el-col :span="1">
											<div style="color: #ff9900; font-size: 16px">{{this.stars2 + "%"}}
											</div>
											</el-col>
										</el-row>

										<el-row>
											<el-col :span="16">
                    <el-rate
											allow-half
                      v-model="v1"
                      disabled
                      text-color="#ff9900"
                      score-template="{value}"
											:colors="['#99A9BF', '#F7BA2A', '#FF9900']"
                    ></el-rate>
										</el-col>
										<el-col :span="1">
											<div style="color: #ff9900; font-size: 16px">{{this.stars1 + "%"}}
											</div>
											</el-col>
										</el-row>

										</div>
										
									</div>
                </el-aside>
                <el-main>
                  <div class="text item">
                  </div>
                  <div class="text item">{{'类型：' + getAllElements(this.currentFilm.genres)}}</div>
									<div class="text item">{{'制片国家/地区：' + getAllElements(this.currentFilm.countries)}}</div>
									<div class="text item">{{'语言：' + getAllElements(this.currentFilm.languages)}}</div>
                  <div class="text item">{{'上映时间：' + this.currentFilm.pubdate}}</div>
									<div class="text item">{{'片长：' + this.currentFilm.duration + '分钟'}}</div>
									<div class="text item">{{'IMDb链接：' + this.currentFilm.imdb}}</div>
                  <div class="text item">{{'导演：' + getAllNames(this.currentFilm.directors)}}</div>
									<div class="text item">{{'编剧：' + getAllNames(this.currentFilm.writers)}}</div>
                  <div class="text item">{{'主演：' + getAllNames(this.currentFilm.casts)}}</div>
                  <div class="text item">{{'剧情：' + this.currentFilm.summary}}</div>
									<el-button style="float: right; padding: 3px 0" type="text" @click="returnList">返回列表</el-button>
                </el-main>
              </el-container>
					</el-card>
				</div>
			</el-main>
      <el-main></el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  name: "FilmDetail",
  data() {
    return {
			activeIndex: "3",
			currentFilm: null,
			averageRating: 2,
			stars5: 0,
			stars4: 0,
			stars3: 0,
			stars2: 0,
			stars1: 0,
			v1: 1,
			v2: 2,
			v3: 3,
			v4: 4,
			v5: 5,
    };
  },
  methods: {
		getData() {
			this.currentFilm = this.$route.params.film
			//this.averageRating = Number(this.currentFilm.rating.average)
			this.stars5 = Number(this.currentFilm.rating.stars[0])
			this.stars4 = Number(this.currentFilm.rating.stars[1])
			this.stars3 = Number(this.currentFilm.rating.stars[2])
			this.stars2 = Number(this.currentFilm.rating.stars[3])
			this.stars1 = Number(this.currentFilm.rating.stars[4])
		},
    handleSelect(key, keyPath) {
      if (keyPath[0] == "1"||"2") {
        this.$router.push(
					{
						name: 'list',
						params: {
							index: Number(keyPath[0])
						}
					}
				);
      }
		},
		returnList() {
			this.$router.push('/')
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
  },
  created() {},
  mounted() {
		this.currentFilm = null
		this.getData()
		console.log(this.currentFilm);
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.nav {
  width: 100%;
}

.image {
  margin-left: 30px;
  margin-top: 25px;
  width: 80%;
  display: block;
}

.stars {
	margin: 20px;
	padding-left: 30px;
}

.title {
  color: #0984e3;
  font-weight: 500;
  font-size: 20px;
}

.text {
  font-size: 14px;
  text-align: left;
}

.item {
  margin-bottom: 15px;
  overflow: hidden;
  display: -webkit-box;
  -webkit-box-orient: vertical;
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
  margin: 15px 150px;

}

</style>
