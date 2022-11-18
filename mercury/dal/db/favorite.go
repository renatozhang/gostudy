package db

import (
	"github.com/renatozhang/gostudy/mercury/common"
	"github.com/renatozhang/gostudy/mercury/logger"
)

func CreateFavoriteDir(favoriteDir *common.FavoriteDir) (err error) {
	// 先查相同的dir_name是否存在
	tx, err := DB.Begin()
	if err != nil {
		logger.Error("Create favorite dir failed, favoriteDir:%#v, err:%v", favoriteDir, err)
		return
	}

	defer func() {
		if err != nil {
			logger.Debug("tx roolback, err:%v", err)
			tx.Rollback()
			return
		}
	}()

	// 根据dir_name查询是否存在
	var dirCount int64
	sqlstr := "select count(dir_id) from favorite_dir where user_id=? and dir_name=?"
	err = DB.Get(&dirCount, sqlstr, favoriteDir.UserId, favoriteDir.DirName)
	if err != nil {
		logger.Error("select dir_name id failed, err:%v, favoriteDir:%#v", err, favoriteDir)
		return
	}

	if dirCount > 0 {
		err = ErrRecordExists
		return
	}

	sqlstr = `insert into favorite_dir(
					user_id,dir_id,dir_name)
			values(?,?,?)`
	_, err = tx.Exec(sqlstr, favoriteDir.UserId, favoriteDir.DirId, favoriteDir.DirName)
	if err != nil {
		logger.Error("create favorite dir failed, favoriteDir:%#v, err:%v", favoriteDir, err)
		return
	}
	err = tx.Commit()
	return
}

func CreateFavorite(favorite *common.Favorite) (err error) {
	tx, err := DB.Begin()
	if err != nil {
		logger.Error("Create favorite dir failed, favorite:%#v, err:%v", favorite, err)
		return
	}

	defer func() {
		if err != nil {
			logger.Debug("tx roolback, err:%v", err)
			tx.Rollback()
			return
		}
	}()

	var favoriteCount int64
	sqlstr := "select count(answer_id) from favorite where user_id=? and dir_id=? and answer_id=?"
	err = DB.Get(&favoriteCount, sqlstr, favorite.UserId, favorite.DirId, favorite.AnswerId)
	if err != nil {
		logger.Error("select dir_name id failed, err:%v, favorite:%#v", err, favorite)
		return
	}
	if favoriteCount > 0 {
		err = ErrRecordExists
		return
	}

	sqlstr = `insert into favorite(
		answer_id,user_id,dir_id)
			values(?,?,?)`
	_, err = tx.Exec(sqlstr, favorite.AnswerId, favorite.UserId, favorite.DirId)
	if err != nil {
		logger.Error("create favorite failed, favorite:%#v, err:%v", favorite, err)
		return
	}
	sqlstr = "update favorite_dir set count=count+1 where dir_id=?"
	_, err = tx.Exec(sqlstr, favorite.DirId)
	if err != nil {
		logger.Error("update favorite dir failed, favorite:%#v, err:%v", favorite, err)
		return
	}
	err = tx.Commit()
	return
}

func GetFavoriteDirList(UserId int64) (favoriteDirList []*common.FavoriteDir, err error) {
	sqlstr := "select dir_id,dir_name,user_id,count from favorite_dir where user_id=?"
	err = DB.Select(&favoriteDirList, sqlstr, UserId)
	if err != nil {
		logger.Error("select favorite dir list failed, err:%v", err)
		return
	}
	return
}

func GetFavoriteList(dirId, UserId, offset, limit int64) (favoriteList []*common.Favorite, err error) {
	sqlstr := "select answer_id,user_id,dir_id from favorite where user_id=? and dir_id=? limit ?,?"
	err = DB.Select(&favoriteList, sqlstr, UserId, dirId, offset, limit)
	if err != nil {
		logger.Error("select favorite list failed, err:%v", err)
		return
	}
	return
}
