package experiment

import "context"

type ExperimentManager struct {
	DB                     orm.DBService                    `inject:"DB"`
	Orm                    *orm.OrmService                  `inject:"OrmService"`
	Config                 brick.Config                     `inject:"config"`
	Logger                 *logging.LoggerService           `inject:"LoggerService"`
}

func (p *ExperimentManager) Save(ctx context.Context, input protocol.SaveInput) (interface{}, error) {
	node := &model.Node{}
	var nodeOutputs []protocol.NodeOutput
	links := &model.Links{}
	metadata := &model.Metadata{}

	for _, Node := range input.Nodes {
		if Node.Status == enums.InitSuccess || Node.Status == enums.Modified {
			Node.Status = enums.InitSuccess
		}

		node.NodeId = Node.ID
		node.Name = Node.Name
		node.Status = Node.Status
		node.ExperimentId = input.ID
		node.Type = Node.Type
		node.X = Node.X
		node.Y = Node.Y

		if err := p.Orm.CreateCtx(ctx, "node", node); err != nil {
			p.Logger.WithField("node", Node.ID).Errorf("failed to create the node : %v", err)
			return nil, nerror.NewCommonError("save node failure: %v", err)
		}
		nodeOutput := protocol.NodeOutput{
			ID:     Node.ID,
			Name:   Node.Name,
			Type:   Node.Type,
			Status: Node.Status,
		}
		nodeOutputs = append(nodeOutputs, nodeOutput)

		for _, value := range Node.Links {
			links.Source = Node.ID
			links.Target = value

			if err := p.Orm.CreateCtx(ctx, "links", links); err != nil {
				p.Logger.WithField("links", links).Errorf("failed to create the links : %v", err)
				return nil, nerror.NewCommonError("save links failure: %v", err)
			}
			links.ID = 0
		}

		for key, value := range Node.Metadata {
			metadata.NodeId = Node.ID
			metadata.Key = key
			metadata.Value = value

			if err := p.Orm.CreateCtx(ctx, "metadata", metadata); err != nil {
				p.Logger.WithField("metadata", metadata).Errorf("failed to create the metadata : %v", err)
				return nil, nerror.NewCommonError("save metadata failure: %v", err)
			}
			metadata.ID = 0
		}
	}

	return &protocol.SaveOutput{
		ID:      input.ID,
		Name:    input.Name,
		Comment: input.Comment,
		Nodes:   nodeOutputs,
	}, nil
}
