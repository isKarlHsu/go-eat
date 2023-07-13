package command

import (
	"eat/app/model"
	"eat/global"
	"github.com/spf13/cobra"
)

var migrationCmd = &cobra.Command{
	Use: "migration",
	Short: "初始化数据库",
	Long: "初始化数据库",
	Run: func(cmd *cobra.Command, args []string) {
		err := global.DB.Migrator().AutoMigrate(
			&model.UserModel{},
			&model.FoodModel{},
		)
		if err != nil {
			global.Logger.Error("生成数据表失败")
			return
		}
		global.Logger.Info("生成数据表成功")
	},
}

func init()  {
	rootCmd.AddCommand(migrationCmd)
}