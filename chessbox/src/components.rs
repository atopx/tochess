use bevy::{prelude::{Resource, Component, States}, time::{Timer, TimerMode}};

#[derive(Resource)]
pub struct GameTimers {
    pub red: Timer,
    pub black: Timer,
}

impl GameTimers {
    pub fn new() -> Self {
        Self {
            red: Timer::from_seconds(5. * 60., TimerMode::Once),
            black: Timer::from_seconds(5. * 60., TimerMode::Once),
        }
    }
}

impl Default for GameTimers {
    fn default() -> Self {
        Self::new()
    }
}

#[derive(Debug, Clone, Copy, Component, PartialEq, Eq)]
pub struct RedTimer;

#[derive(Debug, Clone, Copy, Component, PartialEq, Eq)]
pub struct BlackTimer;


/// 棋子
#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum Pieces {
    /// 车
	R,
    /// 马
	N,
    /// 象
	B,
    /// 士
    A,
    /// 将
    K,
    /// 跑
    C,
    /// 兵
    P,
}

/// 棋子颜色
#[derive(Default, Debug, Clone, Copy, PartialEq, Eq, Hash, States)]
pub enum PieceColor {
	#[default]
	Red,
	Black,
}


/// 棋子位置
#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub struct Position {
	pub row: i8,
	pub col: i8,
}

impl Position {
	pub const fn new(row: i8, col: i8) -> Self {
		Self { row, col }
	}
}

/// Piece 棋子
#[derive(Debug, Clone, Copy, Component, PartialEq, Eq)]
pub struct Piece {
	pub pos: Position,
	pub amount_moved: u32,
	pub piece_type: Pieces,
	pub color: PieceColor,
}
