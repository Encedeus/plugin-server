package service

import (
	"PluginServer/dto"
	"PluginServer/ent"
	"PluginServer/ent/plugin"
	"context"
	"github.com/google/uuid"
)

func CreatePlugin(pluginInfo dto.CreatePluginDTO, ownerId uuid.UUID) error {
	err := Db.Plugin.Create().
		SetName(pluginInfo.Name).
		SetDescription(pluginInfo.Desc).
		SetRepo(pluginInfo.Repo).
		SetHomepage(pluginInfo.Homepage).
		SetOwnerID(ownerId).
		SetContributors(pluginInfo.Contributors).
		Exec(context.Background())

	return err
}

func GetPluginByName(name string) (*ent.Plugin, error) {
	pluginData, err := Db.Plugin.Query().Where(plugin.Name(name)).First(context.Background())

	if err != nil {
		return nil, err
	}

	return pluginData, err
}
